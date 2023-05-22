const wrapper = document.querySelector('.wrapper');
const loginLink = document.querySelector('.login-link');
const registerLink = document.querySelector('.register-link');
const buttonLoginPopup = document.querySelector('.button__login-popup');
const buttonRegisterPopup = document.querySelector('.button__register-popup');
const closeButton = document.querySelector('.ri-close-line');

registerLink.addEventListener('click', () => {
    wrapper.classList.add('active');
});

loginLink.addEventListener('click', () => {
    wrapper.classList.remove('active');
});

buttonLoginPopup.addEventListener('click', () => {
    wrapper.classList.add('login-popup');
});
// buttonRegisterPopup.addEventListener('click', () => {
//     wrapper.classList.add('register-popup');
//     // wrapper.classList.add('active');
// });

closeButton.addEventListener('click', () => {
    wrapper.classList.remove('login-popup');
    wrapper.classList.remove('active');
});

var buttonLikesDislikes = document.querySelectorAll("bulike/dislike");
console.log(buttonLikesDislikes);