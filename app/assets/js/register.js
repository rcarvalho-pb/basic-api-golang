$('#register-user').on('submit', createUser)

function createUser(event){
  event.preventDefault()
  
  if ($('#password').val() != $('#password-confirm').val()) {
    alert("As senhas não coincidem.")
    return
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      Name: $('#name').val(),
      Email: $('#email').val(),
      Nick: $('#nick').val(),
      Password: $('#password').val()
    }
  }).done(function() {
    alert("Usuário cadastrado com sucesso")
  }).fail(function() {
    alert("Erro ao cadastrar usuário")
  })
}