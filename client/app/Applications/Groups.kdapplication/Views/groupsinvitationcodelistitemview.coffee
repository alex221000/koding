class GroupsInvitationCodeListItemView extends KDListItemView

  constructor:(options = {}, data)->
    options.cssClass = 'formline clearfix'
    options.type     = 'invitation-request invitation-code'

    super options, data

    @editButton = new KDButtonView
      style       : 'clean-gray'
      title       : 'Edit'
      icon        : yes
      iconClass   : 'edit'
      callback    : @bound 'showEditModal'

    @shareButton = new KDButtonView
      style       : 'clean-gray'
      title       : 'Share'
      icon        : yes
      iconClass   : 'share'
      callback    : @bound 'showShareModal'

    @statusText    = new KDCustomHTMLView
      partial     : '<span class="icon"></span><span class="title"></span>'
      cssClass    : 'status hidden'

  viewAppended: JView::viewAppended

  showEditModal:->
    {maxUses} = @getData()
    modal = new KDModalViewWithForms
      cssClass                : 'invitation-share-modal'
      title                   : 'Share Invitation Code'
      overlay                 : yes
      width                   : 400
      height                  : 'auto'
      tabs                    :
        forms                 :
          invite              :
            callback          : (formData)=>
              @getData().modifyMultiuse formData, (err)=>
                if err
                  warn err
                  return new KDNotificationView
                    title    : err.message ? 'An error occured! Please try again later.'
                    duration : 2000
                new KDNotificationView
                  title      : 'Invitation code updated!'
                  duration   : 2000
                modal.destroy()
            buttons           :
              Save            :
                itemClass     : KDButtonView
                style         : 'modal-clean-green'
                type          : 'submit'
                loader        :
                  color       : '#444444'
                  diameter    : 12
              Disable         :
                itemClass     : KDButtonView
                style         : 'modal-clean-red'
                loader        :
                  color       : '#ffffff'
                  diameter    : 12
                callback      : -> log 'deactivated'
              Cancel          :
                style         : 'modal-cancel'
                callback      : -> modal.destroy()
            fields            :
              maxUses         :
                itemClass     : KDInputView
                label         : 'Maximum Uses'
                defaultValue  : maxUses
                validate      :
                  rules       :
                    regExp    : /\d+/i
                  messages    :
                    regExp    : 'numbers only please'

  showShareModal:->
    modal = new KDModalViewWithForms
      cssClass               : 'invitation-share-modal'
      title                  : 'Share Invitation Code'
      overlay                : yes
      width                  : 400
      height                 : 'auto'
      tabs                   :
        forms                :
          invite             :
            buttons          :
              Close          :
                itemClass    : KDButtonView
                style        : 'modal-clean-green'
                loader       :
                  color      : '#ffffff'
                  diameter   : 12
                callback     : -> modal.destroy()
            fields           :
              link           :
                itemClass    : KDInputView
                label        : 'Invitation Link'
                defaultValue : @getInvitationUrl()

  getInvitationUrl:->
    {group, code} = @getData()
    slug  = if group and group isnt 'koding' then "#{group}/" else ''
    "https://#{location.host}/#{slug}Invitation/#{code}"

  markDeleted:->
    @statusText.setClass 'deleted'
    @statusText.$('span.title').html 'Deleted'
    @statusText.unsetClass 'hidden'
    @editButton.hide()
    @shareButton.hide()

  pistachio:->
    {code, maxUses, uses} = @getData()
    """
    <section>
      <div class="buttons">{{> @shareButton}} {{> @editButton}}</div>
      {{> @statusText}}
      <div class="details">
        <div class="code">#{code}</div>
        <div class="usage">#{uses} of #{maxUses} used</div>
      </div>
    </section>
    """
