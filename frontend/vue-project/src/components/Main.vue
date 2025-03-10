<template>
    <header class="header">
      <div class="header-container">
        <h1>SoundClowns</h1>
        <nav class="nav">
          <button 
            class="nav-button yandex-auth"
            @click="handleYandexLogin"
          >
            <span class="button-content">
                Войти 
                <img src="@/assets/icon.svg" alt="Иконка" class="logo-image">
            </span>
          </button>
        </nav>
      </div>
    </header>
  
    <div class="main-content">
      <!-- Герой-секция -->
      <section class="hero">
        <div class="hero-wrapper">
          <div class="hero-content">
            <h2>Добро пожаловать в SoundClowns</h2>
            <p>Платформа для создания общих музыкальных комнат для голосования за музыкальынй трек</p>
            <div class="cta-buttons">
              <button class="cta-button" @click="handleYandexLogin">Создать комнату</button>
              <button class="cta-button outline">Узнать больше</button>
            </div>
          </div>
  
          <!-- Слайд-шоу -->
          <div class="slider-container">
            <div 
              class="slide"
              v-for="(slide, index) in slides"
              :key="index"
              :class="{ active: currentSlide === index }"
            >
              <img 
                :src="slide.image" 
                :alt="slide.alt"
                class="slide-image"
              >
            </div>
            <div class="slider-indicators">
              <button
                v-for="(_, index) in slides"
                :key="index"
                :class="{ active: currentSlide === index }"
                @click="currentSlide = index"
              ></button>
            </div>
          </div>
        </div>
        <div class="visual-wave"></div>
      </section>
  
      <!-- Особенности -->
      <section class="features">
        <h3>Почему SoundClowns?</h3>
        <div class="features-grid">
          <div class="feature-card">
            <div class="icon">🎧</div>
            <h4>Эксклюзивные релизы</h4>
            <p>Ранний доступ к новым трекам независимых исполнителей</p>
          </div>
          <div class="feature-card">
            <div class="icon">🎨</div>
            <h4>Креативное сообщество</h4>
            <p>Общайтесь с музыкантами и создавайте коллабы</p>
          </div>
          <div class="feature-card">
            <div class="icon">🔊</div>
            <h4>Hi-Fi звучание</h4>
            <p>Потоковое вещание в качестве lossless</p>
          </div>
        </div>
      </section>
  
      <!-- Лента треков -->
      <section class="track-feed">
        <h3>Сейчас в тренде</h3>
        <div class="track-grid">
          <div class="track-card" v-for="i in 4" :key="i">
            <div class="album-art"></div>
            <div class="track-info">
              <h5>Название трека {{i}}</h5>
              <p>Исполнитель {{i}}</p>
            </div>
          </div>
        </div>
      </section>
    </div>
  </template>
  
  <script setup>
  import { useAuthStore } from '@/stores/auth'
  import { useRouter } from 'vue-router'
  import { ref, onMounted } from 'vue'
  
  // Импорт изображений для слайдов
  import slide1 from '@/assets/slides/slide1.jpg'
  import slide2 from '@/assets/slides/slide2.jpg'
  import slide3 from '@/assets/slides/slide3.jpg'
  
  const authStore = useAuthStore()
  const router = useRouter()
  
  // Логика слайд-шоу
  const slides = ref([
    { image: slide1, alt: 'Музыкальная студия' },
    { image: slide2, alt: 'Концерт' },
    { image: slide3, alt: 'Наушники' }
  ])
  
  const currentSlide = ref(0)
  
  onMounted(() => {
    const interval = setInterval(() => {
      currentSlide.value = (currentSlide.value + 1) % slides.value.length
    }, 5000)
    
    return () => clearInterval(interval)
  })
  
  // Авторизация через Яндекс
  const handleYandexLogin = () => {
    const clientId = import.meta.env.ClientID
    const redirectUri = encodeURIComponent(import.meta.env.VITE_YANDEX_REDIRECT_URI)
    const state = generateRandomString()
    
    window.location.href = `https://oauth.yandex.ru/authorize?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&state=${state}`
  }
  
  const generateRandomString = () => {
    return Array.from(crypto.getRandomValues(new Uint8Array(16)))
      .map(b => b.toString(16).padStart(2, '0'))
      .join('')
  }
  </script>
  
  <style scoped>
  /* Шапка */
  .header {
    background: var(--bg-300);
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    position: fixed;
    width: 100%;
    top: 0;
    z-index: 1000;
  }
  
  .header-container {
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 2rem;
  }
  
  .nav-button {
    padding: 0.7rem 1.2rem;
    border: none;
    border-radius: 8px;
    background: var(--primary-200);
    color: var(--text-200);
    cursor: pointer;
    transition: all 0.3s ease;
  }
  
  .button-content {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-weight: 500;
  }
  
  .logo-image {
    width: 24px;
    height: 24px;
    filter: brightness(0) invert(1);
  }
  
  /* Основной контент */
  .main-content {
    margin-top: 70px;
    padding: 2rem;
  }
  
  /* Герой-секция */
  .hero-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 2rem;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
  }
  
  .hero-content {
    flex: 0 1 50%;
    text-align: left;
  }
  
  .hero h2 {
    font-size: 2.5rem;
    color: var(--primary-300);
    margin-bottom: 1rem;
  }
  
  .hero p {
    font-size: 1.2rem;
    color: var(--text-200);
    margin-bottom: 2rem;
  }
  
  /* Слайд-шоу */
  .slider-container {
    flex: 0 1 45%;
    height: 400px;
    border-radius: 20px;
    overflow: hidden;
    position: relative;
    box-shadow: 0 10px 20px rgba(0,0,0,0.2);
  }
  
  .slide {
    position: absolute;
    width: 100%;
    height: 100%;
    opacity: 0;
    transition: opacity 1s ease;
  }
  
  .slide.active {
    opacity: 1;
  }
  
  .slide-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    filter: grayscale(20%) brightness(0.8);
  }
  
  .slider-indicators {
    position: absolute;
    bottom: 20px;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    gap: 10px;
    z-index: 2;
  }
  
  .slider-indicators button {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    border: none;
    background: rgba(255,255,255,0.3);
    cursor: pointer;
    transition: all 0.3s ease;
  }
  
  .slider-indicators button.active {
    background: var(--accent-200);
    transform: scale(1.2);
  }
  
  /* Особенности */
  .features {
    margin: 4rem 0;
  }
  
  .features h3 {
    text-align: center;
    color: var(--primary-300);
    margin-bottom: 3rem;
  }
  
  .features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
  }
  
  .feature-card {
    background: var(--bg-300);
    padding: 2rem;
    border-radius: 15px;
    text-align: left;
  }
  
  .feature-card .icon {
    font-size: 3rem;
    margin-bottom: 1rem;
    display: inline-block;
  }
  
  /* Лента треков */
  .track-feed {
    margin-top: 4rem;
  }
  
  .track-feed h3 {
    color: var(--primary-300);
    margin-bottom: 2rem;
    text-align: center;
  }
  
  .track-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1.5rem;
    max-width: 1200px;
    margin: 0 auto;
  }
  
  .track-card {
    background: var(--bg-300);
    border-radius: 10px;
    padding: 1rem;
  }
  
  .album-art {
    background: var(--primary-200);
    aspect-ratio: 1;
    border-radius: 8px;
    margin-bottom: 1rem;
  }

    /* Остальные стили */
    .cta-buttons {
    display: flex;
    gap: 1rem;
    margin-top: 2rem;
  }
  
  .cta-button {
    padding: 1rem 2rem;
    border: none;
    border-radius: 50px;
    background: var(--primary-300);
    color: var(--bg-200);
    cursor: pointer;
    transition: transform 0.3s ease;
  }
  
  .cta-button.outline {
    background: none;
    border: 2px solid var(--primary-300);
    color: var(--primary-300);
  }
  
  /* Адаптивность */
  @media (max-width: 768px) {
    .hero-wrapper {
      flex-direction: column;
      padding: 0 1rem;
    }
    
    .slider-container {
      width: 100%;
      order: -1;
      height: 300px;
    }
    
    .feature-card {
      text-align: center;
    }
    
    .feature-card .icon {
      display: block;
      margin: 0 auto 1rem;
    }
    
    .track-grid {
      grid-template-columns: 1fr;
    }
  }
  </style>