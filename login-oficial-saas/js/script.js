'use strict';

/* ==========================================================================
   Utilidades
   ========================================================================== */

const qs = (selector, scope = document) => scope.querySelector(selector);
const qsa = (selector, scope = document) => Array.from(scope.querySelectorAll(selector));

/* ==========================================================================
   Menu mobile (hamburger)
   ========================================================================== */

function initMobileMenu() {
  const menuToggle = qs('#menuToggle');
  const mobileMenu = qs('#mobileMenu');

  if (!menuToggle || !mobileMenu) return;

  const closeMenu = () => {
    menuToggle.setAttribute('aria-expanded', 'false');
    menuToggle.setAttribute('aria-label', 'Abrir menu');
    mobileMenu.classList.remove('mobile-menu--open');
  };

  const openMenu = () => {
    menuToggle.setAttribute('aria-expanded', 'true');
    menuToggle.setAttribute('aria-label', 'Fechar menu');
    mobileMenu.classList.add('mobile-menu--open');
  };

  menuToggle.addEventListener('click', () => {
    const isOpen = menuToggle.getAttribute('aria-expanded') === 'true';
    isOpen ? closeMenu() : openMenu();
  });

  // Fecha o menu ao clicar em um link
  qsa('.mobile-menu__link', mobileMenu).forEach((link) => {
    link.addEventListener('click', closeMenu);
  });

  // Fecha o menu com a tecla Escape
  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') closeMenu();
  });
}

/* ==========================================================================
   Campo de pesquisa expansível
   ========================================================================== */

function initSearch() {
  const search = qs('#search');
  const searchToggle = qs('#searchToggle');
  const searchInput = qs('#searchInput');

  if (!search || !searchToggle || !searchInput) return;

  const openSearch = () => {
    search.classList.add('search--open');
    searchToggle.setAttribute('aria-expanded', 'true');
    searchToggle.setAttribute('aria-label', 'Fechar pesquisa');
    searchInput.focus();
  };

  const closeSearch = () => {
    search.classList.remove('search--open');
    searchToggle.setAttribute('aria-expanded', 'false');
    searchToggle.setAttribute('aria-label', 'Abrir pesquisa');
    searchInput.value = '';
  };

  searchToggle.addEventListener('click', () => {
    const isOpen = search.classList.contains('search--open');
    isOpen ? closeSearch() : openSearch();
  });

  // Fecha ao clicar fora do campo de pesquisa
  document.addEventListener('click', (event) => {
    if (!search.contains(event.target)) closeSearch();
  });

  // Fecha com Escape quando o input está com foco
  searchInput.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') {
      closeSearch();
      searchToggle.focus();
    }
  });
}

/* ==========================================================================
   Mostrar / ocultar senha
   ========================================================================== */

function initPasswordToggle() {
  const toggleButton = qs('#togglePassword');
  const passwordInput = qs('#password');

  if (!toggleButton || !passwordInput) return;

  toggleButton.addEventListener('click', () => {
    const isHidden = passwordInput.getAttribute('type') === 'password';

    passwordInput.setAttribute('type', isHidden ? 'text' : 'password');
    toggleButton.setAttribute('aria-pressed', String(isHidden));
    toggleButton.setAttribute('aria-label', isHidden ? 'Ocultar senha' : 'Mostrar senha');
  });
}

/* ==========================================================================
   Login social (Google)
   ========================================================================== */

function initGoogleLogin() {
  const googleBtn = qs('#googleLoginBtn');

  if (!googleBtn) return;

  googleBtn.addEventListener('click', () => {
    if (googleBtn.classList.contains('btn-social--loading')) return;

    googleBtn.disabled = true;
    googleBtn.classList.add('btn-social--loading');

    // Simula o redirecionamento/autenticação com o Google (sem backend real)
    window.setTimeout(() => {
      googleBtn.disabled = false;
      googleBtn.classList.remove('btn-social--loading');
    }, 1400);
  });
}

/* ==========================================================================
   Validação do formulário de login
   ========================================================================== */

function isValidEmail(value) {
  const pattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return pattern.test(value);
}

function setFieldError(fieldEl, inputEl, errorEl, message) {
  if (message) {
    fieldEl.classList.add('field--has-error');
    inputEl.setAttribute('aria-invalid', 'true');
    errorEl.textContent = message;
    errorEl.classList.add('field__error--visible');
  } else {
    fieldEl.classList.remove('field--has-error');
    inputEl.setAttribute('aria-invalid', 'false');
    errorEl.textContent = '';
    errorEl.classList.remove('field__error--visible');
  }
}

function initLoginForm() {
  const form = qs('#loginForm');
  const emailInput = qs('#email');
  const passwordInput = qs('#password');
  const emailError = qs('#emailError');
  const passwordError = qs('#passwordError');
  const submitBtn = qs('#submitBtn');

  if (!form || !emailInput || !passwordInput || !submitBtn) return;

  const emailField = emailInput.closest('.field');
  const passwordField = passwordInput.closest('.field');

  function validateEmail() {
    const value = emailInput.value.trim();

    if (value === '') {
      setFieldError(emailField, emailInput, emailError, 'E-mail é obrigatório.');
      return false;
    }

    if (!isValidEmail(value)) {
      setFieldError(emailField, emailInput, emailError, 'Digite um e-mail válido.');
      return false;
    }

    setFieldError(emailField, emailInput, emailError, '');
    return true;
  }

  function validatePassword() {
    const value = passwordInput.value;

    if (value === '') {
      setFieldError(passwordField, passwordInput, passwordError, 'Senha é obrigatória.');
      return false;
    }

    setFieldError(passwordField, passwordInput, passwordError, '');
    return true;
  }

  // Revalida ao sair do campo, para feedback imediato
  emailInput.addEventListener('blur', validateEmail);
  passwordInput.addEventListener('blur', validatePassword);

  // Limpa o erro assim que o usuário volta a digitar
  emailInput.addEventListener('input', () => {
    if (emailField.classList.contains('field--has-error')) validateEmail();
  });

  passwordInput.addEventListener('input', () => {
    if (passwordField.classList.contains('field--has-error')) validatePassword();
  });

  function setLoading(isLoading) {
    submitBtn.disabled = isLoading;
    submitBtn.classList.toggle('btn-primary--loading', isLoading);
  }

  form.addEventListener('submit', (event) => {
    event.preventDefault();

    const isEmailValid = validateEmail();
    const isPasswordValid = validatePassword();

    if (!isEmailValid || !isPasswordValid) {
      const firstInvalid = !isEmailValid ? emailInput : passwordInput;
      firstInvalid.focus();
      return;
    }

    setLoading(true);

    // Simula uma chamada de autenticação (sem backend real)
    window.setTimeout(() => {
      setLoading(false);
      submitBtn.blur();
    }, 1400);
  });
}

/* ==========================================================================
   Inicialização
   ========================================================================== */

document.addEventListener('DOMContentLoaded', () => {
  initMobileMenu();
  initSearch();
  initGoogleLogin();
  initPasswordToggle();
  initLoginForm();
});
