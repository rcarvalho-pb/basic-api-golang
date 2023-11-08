const USER_RESOURCE = "/users"

$('#register-user').on('submit', createUser)

function createUser(event){
  event.preventDefault()
  
  if ($('#password').val() != $('#password-confirm').val()) {
    alert("As senhas não coincidem.")
    return
  }

  $.ajax({
    url: USER_RESOURCE,
    method: "POST",
    data: {
      Name: $('#name').val(),
      Email: $('#email').val(),
      Nick: $('#nick').val(),
      Password: $('#password').val()
    }
  }).done(function() {
    Swal.fire(
      'Sucesso!',
      'Usuário registrado com sucesso',
      'success',
    ).then(function() {
      $.ajax({
        url: "/login",
        method: "POST",
        data: {
          email: $('#email').val(),
          password: $("#password").val()
        }
      }).done(function() {
        window.location = '/home'
      }).fail(function() {
        Swal.fire(
          'Ops...!',
          'Erro ao autenticar usuário',
          'error',
        )
      })
    })
  }).fail(function(err) {
    Swal.fire(
      'Ops...!',
      'Erro ao registrar Usuário',
      'error',
    )
  })
}