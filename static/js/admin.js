$('.menu-link').click(function() {
  var speed = 200;
  var link = $(this).attr("href");
  
  $('.menu-link.active').removeClass('active');
  $(this).addClass('active');
  
  $('.content.box-active').animate({opacity: 0}, speed, function() {
    $(this).removeClass('box-active');
    $(link).animate({opacity: 1}, speed).addClass('box-active');
  });
});