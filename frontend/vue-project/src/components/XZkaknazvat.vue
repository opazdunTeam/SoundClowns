<template>
  <div>
    <!-- Шапка -->
    <header class="header">
      <div class="header-container">
        <h1>SoundClowns</h1>
        <nav class="nav">
          <button 
            class="nav-button yandex-auth"
            @click="isAuthenticated ? logout() : loginWithYandex()"
          >
            <span class="button-content">
              {{ isAuthenticated ? 'Выйти' : 'Войти' }}
              <img v-if="!isAuthenticated" src="@/assets/icon.svg" alt="Иконка" class="logo-image">
            </span>
          </button>
        </nav>
      </div>
    </header>

    <main class="main-content">
      <!-- Блок для созданной комнаты -->
      <div v-if="createdRoom" class="created-room">
        <div class="room-card large">
          <div class="card-image">
            <img :src="createdRoom.image" alt="Room cover" class="room-cover">
          </div>
          <div class="card-content">
            <h3 class="room-title">{{ createdRoom.title }}</h3>
            <p class="room-description">{{ createdRoom.description }}</p>
            <div class="room-info">
              <span class="listeners">
                👥 {{ createdRoom.listeners }} слушателей
              </span>
              <button class="delete-button" @click.stop="handleDelete">
                Удалить
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Кнопка создания -->
      <button 
        v-else 
        class="create-room-button" 
        @click="showModal = true"
      >
        Создать комнату
      </button>

      <!-- Список всех комнат -->
      <div class="rooms-list">
        <div 
          v-for="(room, index) in rooms" 
          :key="index"
          class="room-card"
          @click="handleRoomClick(room)"
        >
          <div class="card-image">
            <img :src="room.image" alt="Room cover" class="room-cover">
          </div>
          <div class="card-content">
            <h3 class="room-title">{{ room.title }}</h3>
            <p class="room-description">{{ room.description }}</p>
            <div class="room-info">
              <span class="listeners">
                👥 {{ room.listeners }} слушателей
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Модальное окно -->
      <div v-if="showModal" class="modal">
        <div class="modal-content">
          <span class="close" @click="showModal = false">&times;</span>
          <h2>Создать новую комнату</h2>
          <form @submit.prevent="createRoom">
            <div class="form-group">
              <label>Название:</label>
              <input v-model="newRoom.title" required>
            </div>
            <div class="form-group">
              <label>Описание:</label>
              <textarea v-model="newRoom.description" required></textarea>
            </div>
            <div class="form-group">
              <label>Ссылка на изображение:</label>
              <input v-model="newRoom.image" required type="url">
            </div>
            <button type="submit" class="submit-button">Создать</button>
          </form>
        </div>
      </div>
    </main>
  </div>
</template>


<script>
export default {
  data() {
    return {
      showModal: false,
      createdRoom: null,
      newRoom: {
        title: '',
        description: '',
        image: '',
        listeners: 0
      },
      rooms: [
        {
          id: 1,
          title: 'Рок хиты 90-х',
          description: 'Лучшие рок-композиции из 90-х годов',
          listeners: 256,
          image: 'https://artchive.ru/res/media/img/orig/article/89b/807207.webp'
        },
        {
          id: 2,
          title: 'Электронная волна',
          description: 'Свежие треки из мира электронной музыки',
          listeners: 184,
          image: 'https://artchive.ru/res/media/img/orig/article/89b/807207.webp'
        }
      ],
      isAuthenticated: false
    }
  },
  mounted() {
    this.checkAuthState()
    
    const queryParams = this.$route.query
    if (queryParams.state) {
      const authData = {
        state: queryParams.state,
        code: queryParams.code,
        cid: queryParams.cid,
        timestamp: new Date().getTime()
      }
      
      localStorage.setItem('yandexAuth', JSON.stringify(authData))
      this.isAuthenticated = true
      this.$router.replace({ path: this.$route.path, query: {} })
    }
  },
  methods: {
    checkAuthState() {
      const authData = localStorage.getItem('yandexAuth')
      this.isAuthenticated = !!authData
    },
    loginWithYandex() {
      const clientId = import.meta.env.VITE_CLIENT_ID
      const redirectUri = encodeURIComponent(import.meta.env.VITE_YANDEX_REDIRECT_URI)
      const state = this.generateRandomString()
      
      localStorage.setItem('yandexAuthState', state)
      
      window.location.href = `https://oauth.yandex.ru/authorize?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&state=${state}`
    },
    logout() {
      localStorage.removeItem('yandexAuth')
      this.isAuthenticated = false
      this.$router.push('/')
    },
    generateRandomString() {
      return Array.from(crypto.getRandomValues(new Uint8Array(16)))
        .map(b => b.toString(16).padStart(2, '0'))
        .join('')
    },
    handleDelete() {
      this.createdRoom = null
    },
    createRoom() {
      this.createdRoom = {
        ...this.newRoom,
        id: Date.now(),
        listeners: 0
      }
      this.showModal = false
      this.newRoom = {
        title: '',
        description: '',
        image: '',
        listeners: 0
      }
    },
    handleRoomClick(room) {
      this.$router.push(`/rooms/${room.id}`)
    }
  }
}
</script>


<style scoped>
/* Цветовая схема */
:root {
  --bg-200: #292e3b;
  --bg-300: #414654;
  --primary-200: #56647b;
  --primary-300: #b4c2dc;
  --accent-200: #ffecda;
  --text-200: #e0e0e0;
  --bg-300-rgb: 65, 70, 84;
  --primary-300-rgb: 180, 194, 220;
}

/* Основные стили */
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

.main-content {
  margin-top: 80px;
  padding: 2rem;
}

.create-room-button {
  background: var(--primary-200);
  color: white;
  border: none;
  padding: 1rem 2rem;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.1rem;
  display: block;
  width: 100%;
  max-width: 800px;
  margin: 0 auto 2rem;
  transition: transform 0.2s;
}

.create-room-button:hover {
  transform: scale(1.02);
}

.rooms-list {
  max-width: 800px;
  margin: 0 auto;
}

.room-card {
  background: var(--bg-300);
  border: 1px solid var(--primary-200);
  margin-bottom: 1.5rem;
  display: flex;
  min-height: 150px;
  transition: all 0.3s ease;
  cursor: pointer;
  color: var(--text-200);
  text-decoration: none;
}

.room-card:hover {
  background: var(--bg-200);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.card-image {
  width: 200px;
  flex-shrink: 0;
}

.room-cover {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.card-content {
  padding: 1.5rem;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.room-title {
  color: var(--text-200);
  margin: 0 0 0.5rem;
  font-size: 1.25rem;
}

.room-description {
  color: var(--text-200);
  font-size: 0.9rem;
  margin: 0;
  line-height: 1.4;
}

.room-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 1rem;
}

.listeners {
  color: var(--primary-300);
  font-size: 0.85rem;
}

.delete-button {
  background: #ff4444;
  color: white;
  border: none;
  padding: 0.8rem 2rem;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  transition: all 0.3s ease;
}

.delete-button:hover {
  background: #cc0000;
  transform: translateY(-2px);
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1001;
}

.modal-content {
  background: var(--bg-300);
  padding: 2rem;
  border-radius: 10px;
  width: 90%;
  max-width: 500px;
  position: relative;
  color: var(--text-200);
}

.close {
  position: absolute;
  right: 20px;
  top: 10px;
  font-size: 28px;
  cursor: pointer;
  color: var(--text-200);
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: var(--text-200);
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.8rem;
  background: var(--bg-200);
  border: 1px solid var(--primary-200);
  border-radius: 4px;
  font-size: 1rem;
  color: var(--text-200);
}

.submit-button {
  background: var(--primary-200);
  color: var(--text-200);
  padding: 1rem 2rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  width: 100%;
  font-size: 1.1rem;
  margin-top: 1rem;
  transition: all 0.3s ease;
}

.submit-button:hover {
  background: var(--primary-300);
}

@media (max-width: 600px) {
  .room-card {
    flex-direction: column;
  }
  
  .card-image {
    width: 100%;
    height: 200px;
  }
  
  .main-content {
    padding: 1rem;
  }
  
  .modal-content {
    padding: 1rem;
  }
}
</style>