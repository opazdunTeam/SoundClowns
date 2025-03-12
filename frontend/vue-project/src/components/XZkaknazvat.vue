<template>
    <header class="header">
      <div class="header-container">
        <h1>SoundClowns</h1>
        <nav class="nav">
          <button class="nav-button yandex-auth">
            <span class="button-content">
              Войти
              <img src="@/assets/icon.svg" alt="Иконка" class="logo-image">
            </span>
          </button>
        </nav>
      </div>
    </header>
  
    <main class="main-content">
      <!-- Блок для созданной комнаты -->
      <div v-if="createdRoom" class="created-room">
        <router-link 
          :to="`/rooms/${createdRoom.id}`" 
          class="room-card large"
        >
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
        </router-link>
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
        <router-link 
          v-for="(room, index) in rooms" 
          :key="index"
          :to="`/rooms/${room.id}`"
          class="room-card"
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
        </router-link>
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
            image: 'https://example.com/rock-cover.jpg'
          },
          {
            id: 2,
            title: 'Электронная волна',
            description: 'Свежие треки из мира электронной музыки',
            listeners: 184,
            image: 'https://example.com/electronic-cover.jpg'
          }
        ]
      }
    },
    methods: {
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
      }
    }
  }
  </script>

  
  <style scoped>
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
    background: var(--bg-100);
    border: 1px solid var(--border-color);
    margin-bottom: 1.5rem;
    display: flex;
    min-height: 150px;
    transition: all 0.3s ease;
    cursor: pointer;
  }
  
  .room-card:hover {
    background: var(--bg-200);
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  }
  
  .new-room {
    transform: scale(1.05);
    border: 2px solid var(--primary-200);
    margin: 2rem 0;
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
    color: var(--text-100);
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
    color: var(--text-300);
    font-size: 0.85rem;
  }
  
  .join-button {
    background: var(--primary-300);
    color: white;
    border: none;
    padding: 0.8rem 2rem;
    border-radius: 6px;
    cursor: pointer;
    font-weight: bold;
    transition: all 0.3s ease;
  }
  
  .join-button:hover {
    background: var(--primary-400);
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
    background: white;
    padding: 2rem;
    border-radius: 10px;
    width: 90%;
    max-width: 500px;
    position: relative;
  }
  
  .close {
    position: absolute;
    right: 20px;
    top: 10px;
    font-size: 28px;
    cursor: pointer;
  }
  
  .form-group {
    margin-bottom: 1.5rem;
  }
  
  .form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
  }
  
  .form-group input,
  .form-group textarea {
    width: 100%;
    padding: 0.8rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;
  }
  
  .submit-button {
    background: var(--primary-200);
    color: white;
    padding: 1rem 2rem;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    width: 100%;
    font-size: 1.1rem;
    margin-top: 1rem;
  }

  .create-room-button {
  background: var(--primary-200);
  color: white;
  border: none;
  padding: 1.5rem 2rem;
  border-radius: 12px;
  cursor: pointer;
  margin: 2rem auto;
  display: block;
  width: 100%;
  max-width: 800px;
  font-size: 1.2rem;
  transition: all 0.3s ease;
}

.created-room {
  max-width: 800px;
  margin: 0 auto;
}

.room-card.large {
  margin: 2rem 0;
  min-height: 300px;
  border-radius: 15px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
}

.room-card.large .card-image {
  width: 300px;
  height: 300px;
}

.room-card.large .card-content {
  padding: 2rem;
}

.room-card.large .room-title {
  font-size: 1.5rem;
}

.room-card.large .room-description {
  font-size: 1.1rem;
}

.rooms-list {
  margin-top: 3rem;
}

/* Новые стили для кнопки удаления */
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

/* Убираем отступ снизу для комнат в списке */
.rooms-list .room-card {
  margin-bottom: 1.5rem;
}

/* Убираем ховер-эффект для комнат в списке */
.rooms-list .room-card:hover {
  background: var(--bg-100);
  box-shadow: none;
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