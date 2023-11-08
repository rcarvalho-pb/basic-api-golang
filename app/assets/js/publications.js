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
    Swal.fire(
      'Sucesso!',
      'Erro ao criar publicação!',
      'error',
    )
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
    Swal.fire(
      'Sucesso!',
      'Erro ao curtir publicação!',
      'error',
    )
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
    Swal.fire(
      'Sucesso!',
      'Erro ao descurtir publicacao!',
      'error',
    )
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
    Swal.fire(
      'Sucesso!',
      'Atualização realizada com sucesso',
      'success',
    ).then(function() {
      window.location = '/home'
    })
  }).fail(function() {
    Swal.fire({
      icon: 'error',
      title: 'Erro ao realizar atualização da publicação',
      showConfirmButton: false,
      timer: 1500
    })
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

   
  const swalWithBootstrapButtons = Swal.mixin({
    customClass: {
      confirmButton: 'btn btn-success',
      cancelButton: 'btn btn-danger'
    },
    buttonsStyling: false
  })
  
  Swal.fire({
    title: 'Tem certeza que deseja deletar?',
    text: "Não é possível desfazer após confirmar",
    icon: 'warning',
    showCancelButton: true,
    cancelButtonText: 'Não, cancelar!',
    confirmButtonText: 'Deletar!'
  }).then((result) => {
    if (result.isConfirmed) {
      $.ajax({
        url: `${PUBLICATION_RESOURCE}/${publicationId}`,
        method: "DELETE",
      }).done(function() {
        swalWithBootstrapButtons.fire(
          'Deleteda!',
          'Sua publicação foi deletada.',
          'success'
        )
        publication.fadeOut("slow", function() {
          $(this).remove()
        })
      }).fail(function() {
        Swal.fire({
          icon: 'error',
          title: 'Erro ao excluir a publicação',
          showConfirmButton: false,
          timer: 1500
        })
      })
    } else if (
      /* Read more about handling dismissals below */
      result.dismiss === Swal.DismissReason.cancel
    ) {
      swalWithBootstrapButtons.fire(
        'Cancelled',
        'Your imaginary file is safe :)',
        'error'
      )
    }
  })
}