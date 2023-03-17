let select = document.querySelector('select');

select.addEventListener('focus', () => {
  select.size = 5; 
  select.classList.add('fadeIn'); 
  select.classList.remove('fadeOut');
  select.style.backgroundColor = '#FFF';
});

select.addEventListener('blur', () => {
  select.size = 1; 
  select.classList.add('fadeOut');
  select.classList.remove('fadeIn');
  select.style.backgroundColor = '#FFF';
});

select.addEventListener('change', () => {
  select.size = 1; 
  select.blur();
  select.style.backgroundColor = '#FFF';
});

select.addEventListener('mouseover', () => {
  if(select.size == 1){
     select.style.backgroundColor = 'rgb(247, 247, 247)';
  }
});
select.addEventListener('mouseout', () => {
  if(select.size == 1){
     select.style.backgroundColor = '#FFF';
  }
});