class JInvitationRequest extends bongo.Model

  @share()
  
  @set
    indexes       :
      email       : 'unique'
    sharedMethods :
      static      : ['create']
    schema        :
      email       :
        type      : String
        email     : yes
      requestedAt :
        type      : Date
        default   : -> new Date

  @create =({email}, callback)->
    invite = new @ {email}
    invite.save (err)->
      if err
        callback err
      else
        callback null, email