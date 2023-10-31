const PUBLICATION_RESOURCE = "/publications"

$("#new-publication").on('submit', createPublication)
$(document).on('click', '.like-publication', likePublication)
$(document).on('click', '.dislike-publication', dislikePublication)
$("#update-publication").on('click', updatePublication)
$(".delete-publication").on('click', deletePublication)

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
  event.preventDefault()

  const element = $(event.target)

  const publicationId = element.closest('div').data('id-publication')

  element.prop('disable', true)

  $.ajax({
    url: `${PUBLICATION_RESOURCE}/${publicationId}/like`,
    method: "PATCH"
  }).done(function (){
    const likesCounter = element.next('span')
    const likes = parseInt(likesCounter.text())

    likesCounter.text(likes + 1)

    element.addClass('dislike-publication')
    element.addClass('text-danger')
    element.removeClass('like-publication')

  }).fail(function (err){
    console.log(err)
    alert("Erro ao curtir publicacao")
  }).always(function() {
    element.prop('disable', false)
  })
}

function dislikePublication(event) {
  event.preventDefault()

  const element = $(event.target)

  const publicationId = element.closest('div').data('id-publication')

  element.prop('disable', true)

  $.ajax({
    url: `${PUBLICATION_RESOURCE}/${publicationId}/dislike`,
    method: "PATCH"
  }).done(function (){
    const likesCounter = element.next('span')
    const likes = parseInt(likesCounter.text())

    likesCounter.text(likes - 1)

    element.removeClass('dislike-publication')
    element.removeClass('text-danger')
    element.addClass('like-publication')

  }).fail(function (err){
    console.log(err)
    alert("Erro ao descurtir publicacao")
  }).always(function() {
    element.prop('disable', false)
  })

}

function updatePublication() {
  $(this).prop('disable', true)

  const publicationId = $(this).data("publication-id")

  $.ajax({
    url: `${PUBLICATION_RESOURCE}/${publicationId}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val()
    }
  }).done(function() {
    alert("publicação editada com sucesso")
  }).fail(function() {
    alert("erro ao editar publicação")
  }).always(
    $("#update-publication").prop("disable", false)
  )
}

function deletePublication(event) {
  event.preventDefault()

  const element = $(event.target)

  const publication = element.closest('div')
  const publicationId = publication.data('id-publication')

  element.prop('disable', true)

  $.ajax({
    url: `${PUBLICATION_RESOURCE}/${publicationId}`,
    method: "DELETE",
  }).done(function() {
    publication.fadeOut("slow", function() {
      $(this).remove()
    })
  }).fail(function() {
    alert("Erro ao excluir publicacao")
  })
}