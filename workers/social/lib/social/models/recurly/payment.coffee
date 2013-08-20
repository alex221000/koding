jraphical = require 'jraphical'
recurly   = require 'koding-payment'

module.exports = class JRecurlyPayment extends jraphical.Module

  {secure, daisy}      = require 'bongo'

  KodingError          = require '../../error'
  JUser                = require '../user'
  JGroup               = require '../group'
  JRecurlyPlan         = require './index'
  JRecurlySubscription = require './subscription'
  {Relationship}       = require 'jraphical'

  @share()

  @set
    sharedMethods  :
      static       : ['makePayment']
      instance     : ['info']
    schema           :
      planCode       : String
      planQuantity   : Number
      buyer          : String # Recurly account name
      user           : String # Recurly account name
      amount         : Number
      timestamp      : Number
      subscription   : String # Recurly UUID for subscription
      active         : Boolean

  @makePayment = secure (client, data, callback)->
    @initialize client, data, (err, group, account, plan)=>
      if err then return callback new KodingError 'Unable to access payment backend, please try again later.'

      data.quantity ?= 1  # Default quantity is 1
      data.multiple ?= no # Don't allow multiple subscriptions by default

      user  = buyer = "user_#{account._id}"
      buyer = "group_#{group._id}"  if data.chargeTo is 'group'

      charge = (subscription)->
        pay = new JRecurlyPayment {
          buyer
          user
          subscription
          planCode     : data.plan
          planQuantity : data.quantity
          amount       : 0
          active       : yes
        }
        pay.save (err)->
          if err then return callback new KodingError 'Unable to save transaction to database.'
          callback null, pay

      if data.chargeTo is 'group'
        @canChargeGroup group, account, data, (err)=>
          return callback err  if err
          @chargeGroup group, plan, data, (err, subscription)=>
            return callback err  if err
            charge subscription.uuid
      else
        @chargeUser account, plan, data, (err, subscription)=>
          return callback err  if err
          charge subscription.uuid

  # Get ready for a transaction (find group, create accounts)
  @initialize = (client, data, callback)->
    account = client.connection.delegate
    slug    = client.context.group

    JGroup.one {slug}, (err, group)=>
      return callback err  if err
      @createAccount account, (err)=>
        return callback err  if err
        @createGroupAccount group, (err)=>
          return callback err  if err
          @getPlan data.plan, (err, plan)->
            return callback err  if err
            callback null, group, account, plan

  # Get plan
  @getPlan = (code, callback)->
    JRecurlyPlan.getPlanWithCode code, (err, plan)->
      if err then return callback new KodingError 'Unable to access product information. Please try again later.'
      callback null, plan

  # Get price for a product
  @calculateAmount = (data, callback)->
    @getPlan data.plan, (err, plan)->
      return callback err  if err
      callback null, plan.feeMonthly * data.quantity

  # Create user account on Recurly
  @createAccount = (account, callback)->
    account.fetchUser (err, user) ->
      return callback err  if err
      {nickname, firstName, lastName} = account.profile
      data = {email: user.email, nickname, firstName, lastName}
      recurly.setAccount "user_#{account._id}", data, callback

  # Create group account on Recurly
  @createGroupAccount = (group, callback)->
    group.fetchOwner (err, owner)->
      return callback err  if err
      owner.fetchUser (err, user)->
        return callback err  if err
        data =
          username  : group.slug
          firstName : 'Group'
          lastName  : group.title
          email     : user.email
        recurly.setAccount "group_#{group._id}", data, callback

  # Tell if user can buy an item and expense it to group.
  @canChargeGroup = (group, account, data, callback)->
    @calculateAmount data, (err, amount)=>
      return callback err  if err
      @getGroupAllocation group, (err, allocation)=>
        return callback err  if err
        @getUserExpenses group, account, (err, expenses)=>
          return callback err  if err
          if allocation >= expenses + amount
            callback()
          else
            callback new KodingError "You don't have enough balance."

  # Return quota that group gives to its users.
  @getGroupAllocation = (group, callback)->
    group.fetchBundle (err, bundle)=>
      if err or not bundle or bundle.allocation is 0
        return callback new KodingError "This group doesn't allow you to purchase."
      else
        callback null, bundle.allocation

  # Charge group account
  @chargeGroup = (group, p, data, callback)->
    {multiple, plan, quantity} = data
    p.subscribeGroup group, {multiple, plan, quantity}, (err, subscription)->
      if err then return callback new KodingError "Unable to buy item: #{err}"
      callback null, subscription

  # Charge user account
  @chargeUser = (account, plan, data, callback)->
    return callback new KodingError 'Unable charge, insufficient funds.'
    # TBI

  # List group's payments
  @getGroupPayments = (group, callback)->
    @getExpenses
      buyer     : "group_#{group._id}"
      active    : yes
    , callback

  # List user's payments
  @getUserPayments = (account, callback)->
    @getExpenses
      buyer     : "user_#{account._id}"
      active    : yes
    , callback

  # List a user's payments expensed to group
  @getUserExpenses = (group, account, callback)->
    @getExpenses
      buyer     : "group_#{group._id}"
      user      : "user_#{account._id}"
      active    : yes
    , callback

  # TODO: Make sure this calculation is enough.
  #       Not tested for expired/canceled subscriptions.
  @getExpenses = (pattern, callback)->
    error = (err)-> new KodingError "Unable to query user balance: #{err}"

    JRecurlyPayment.some pattern, {subscription: 1}, (err, items)->
      return callback error err  if err
      recurly.getSubscriptions "group_#{group._id}", (err, subs)->
        return callback error err  if err
        uuids = (item.subscription for item in items)
        expenses = 0
        for sub in subs when sub.uuid in uuids
          expenses += parseInt sub.amount, 10

        callback null, expenses
