$('#login').on('submit', login)

function login(event) {
  event.preventDefault()

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      password: $("#password").val()
    }
  }).done(function() {
    alert("Sucesso ao logar")
    window.location = "/huome"
  }).fail(function() {
    alert("Usuário inválido")
  });
}