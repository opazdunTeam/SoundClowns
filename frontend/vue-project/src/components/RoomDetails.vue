<template>
  <div class="room-container">
    <!-- Шапка комнаты -->
    <header class="room-header">
      <button class="back-button" @click="$router.go(-1)">
        <span class="back-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
            <path d="M15 18L9 12L15 6" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </span>
        <span class="back-text">Назад</span>
      </button>
      <h1 class="room-title">{{ roomName }}</h1>
    </header>

    <div class="room-content">
      <!-- Плейлист -->
      <div class="playlist-section glass-card">
        <h2 class="section-title">🎵 Текущий плейлист</h2>
        <SongList :songs="approvedSongs" :show-actions="false"/>
      </div>

      <!-- Предложенные песни -->
      <div class="requests-section glass-card">
        <h2 class="section-title">
          <span v-if="isCreator">📥 Заявки на добавление</span>
          <span v-else>🎧 Предложенные треки</span>
        </h2>
        
        <SongList 
          :songs="pendingSongs" 
          :show-actions="isCreator"
          @approve="handleApprove"
          @reject="handleReject"
        />

        <!-- Форма предложения песни -->
        <div v-if="!isCreator" class="song-request-form">
          <div class="input-group">
            <input 
              v-model="newSongTitle" 
              placeholder="Название трека"
              class="styled-input"
            />
            <span class="input-icon">🎶</span>
          </div>
          
          <div class="input-group">
            <input 
              v-model="newSongArtist" 
              placeholder="Исполнитель"
              class="styled-input"
            />
            <span class="input-icon">🎤</span>
          </div>

          <button 
            @click="submitSongRequest" 
            class="submit-button"
            :disabled="!isFormValid"
          >
            <span class="button-content">
              ✨ Предложить трек
            </span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SongList from './SongList.vue'

export default {
  components: { SongList },
  props: {
    isCreator: Boolean,
    roomName: String
  },
  data() {
    return {
      newSongTitle: '',
      newSongArtist: '',
      approvedSongs: [
        { id: 1, title: 'Bohemian Rhapsody', artist: 'Queen', duration: '5:55' }
      ],
      pendingSongs: [
        { id: 2, title: 'Hotel California', artist: 'Eagles', duration: '6:30', requester: 'Гость 1' }
      ]
    }
  },
  computed: {
    isFormValid() {
      return this.newSongTitle.trim() && this.newSongArtist.trim()
    }
  },
  methods: {
    submitSongRequest() {
      if (this.isFormValid) {
        this.pendingSongs.push({
          id: Date.now(),
          title: this.newSongTitle,
          artist: this.newSongArtist,
          duration: '0:00',
          requester: 'Вы'
        })
        this.resetForm()
      }
    },
    handleApprove(songId) {
      const songIndex = this.pendingSongs.findIndex(s => s.id === songId)
      const [approvedSong] = this.pendingSongs.splice(songIndex, 1)
      this.approvedSongs.push(approvedSong)
    },
    handleReject(songId) {
      this.pendingSongs = this.pendingSongs.filter(s => s.id !== songId)
    },
    resetForm() {
      this.newSongTitle = ''
      this.newSongArtist = ''
    }
  }
}
</script>

<style scoped>
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

@media (prefers-color-scheme: dark) {
  :root {
    --bg-200: #1a1e28;
    --bg-300: #2d323f;
    --primary-200: #3d4a63;
    --primary-300: #8a9bbf;
    --accent-200: #ffd8b2;
    --text-200: #d0d0d0;
    --bg-300-rgb: 45, 50, 63;
    --primary-300-rgb: 138, 155, 191;
  }
}

.room-container {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
  min-height: 100vh;
  background: var(--bg-200);
}

.room-header {
  background: var(--bg-300);
  padding: 1rem 2rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  position: sticky;
  top: 0;
  z-index: 100;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  border-bottom: 1px solid var(--primary-200);
}

.back-button {
  background: none;
  border: none;
  color: var(--primary-300);
  font-size: 1.1rem;
  cursor: pointer;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.back-button:hover {
  background: rgba(var(--primary-300-rgb), 0.1);
  transform: translateX(-3px);
}

.back-icon {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.back-icon svg {
  width: 100%;
  height: 100%;
  stroke: var(--primary-300);
}

.room-title {
  font-size: 1.4rem;
  color: var(--text-200);
  font-weight: 600;
}

.room-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin-top: 2rem;
  padding-top: 2rem;
}

.glass-card {
  background: rgba(var(--bg-300-rgb), 0.95);
  border-radius: 20px;
  padding: 2rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
  border: 1px solid var(--primary-200);
}

.section-title {
  font-size: 1.5rem;
  color: var(--primary-300);
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 2px solid var(--primary-200);
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.song-request-form {
  margin-top: 2rem;
  display: flex;
  flex-direction: column;
  gap: 1.2rem;
}

.input-group {
  position: relative;
}

.styled-input {
  width: 100%;
  padding: 1rem 1rem 1rem 3rem;
  background: var(--bg-200);
  border: 2px solid var(--primary-200);
  border-radius: 12px;
  font-size: 1rem;
  color: var(--text-200);
  transition: all 0.3s ease;
}

.styled-input::placeholder {
  color: var(--primary-300);
}

.styled-input:focus {
  outline: none;
  border-color: var(--primary-300);
  box-shadow: 0 0 0 3px rgba(var(--primary-300-rgb), 0.1);
}

.input-icon {
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
  font-size: 1.2rem;
  color: var(--primary-300);
}

.submit-button {
  background: linear-gradient(
    135deg, 
    var(--primary-200) 0%, 
    var(--primary-300) 100%
  );
  color: var(--accent-200);
  padding: 1rem 2rem;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.submit-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.submit-button:not(:disabled):hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(var(--primary-300-rgb), 0.3);
}

@media (max-width: 768px) {
  .room-content {
    grid-template-columns: 1fr;
  }
  
  .glass-card {
    padding: 1.5rem;
  }
  
  .room-header {
    padding: 1rem;
    gap: 1rem;
  }
  
  .back-text {
    display: none;
  }
  
  .room-title {
    font-size: 1.2rem;
  }
}
</style>