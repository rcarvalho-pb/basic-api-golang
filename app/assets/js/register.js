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
    alert("Usuário cadastrado com sucesso")
    window.location = "/"
  }).fail(function(err) {
    console.log(err)
    alert("Erro ao cadastrar usuário")
  })
}