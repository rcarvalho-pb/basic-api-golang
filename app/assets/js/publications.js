$("#new-publication").on('submit', createPublication)
$(".like-publication").on('click', likePublication)

function createPublication(event) {
  event.preventDefault()

  $.ajax({
    url: "/publications",
    method: "POST",
    data: {
      Title: $("#title").val(),
      Content: $("#content").val()
    }
  }).done(function() {
    window.location = "/home"
  }).fail(function() {
    alert("Erro ao criar publicação!")
  })
}

function likePublication(event) {
  
}