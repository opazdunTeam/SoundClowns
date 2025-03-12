<template>
    <div class="song-list">
      <transition-group name="list">
        <div 
          v-for="song in songs" 
          :key="song.id" 
          class="song-item"
        >
          <div class="song-info">
            <div class="title">{{ song.title }}</div>
            <div class="meta">
              <span class="artist">{{ song.artist }}</span>
              <span class="duration">{{ song.duration }}</span>
            </div>
            <div v-if="song.requester" class="requester">
              👤 {{ song.requester }}
            </div>
          </div>
          
          <div v-if="showActions" class="song-actions">
            <button 
              @click="$emit('approve', song.id)" 
              class="action-btn approve"
            >
              <span class="emoji">✅</span>
            </button>
            <button 
              @click="$emit('reject', song.id)" 
              class="action-btn reject"
            >
              <span class="emoji">❌</span>
            </button>
          </div>
        </div>
      </transition-group>
    </div>
  </template>
  
  <style scoped>
  .song-list {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }
  
  .song-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1.2rem;
    background: white;
    border-radius: 12px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
    transition: all 0.3s ease;
  }
  
  .song-item:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }
  
  .title {
    font-weight: 600;
    color: #2c3e50;
    margin-bottom: 0.3rem;
  }
  
  .meta {
    display: flex;
    gap: 1rem;
    color: #7f8c8d;
    font-size: 0.9rem;
  }
  
  .requester {
    margin-top: 0.5rem;
    font-size: 0.85rem;
    color: #6c5ce7;
  }
  
  .song-actions {
    display: flex;
    gap: 0.5rem;
  }
  
  .action-btn {
    width: 40px;
    height: 40px;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.2s ease;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .approve {
    background: rgba(76, 175, 80, 0.1);
  }
  
  .approve:hover {
    background: rgba(76, 175, 80, 0.2);
  }
  
  .reject {
    background: rgba(244, 67, 54, 0.1);
  }
  
  .reject:hover {
    background: rgba(244, 67, 54, 0.2);
  }
  
  .emoji {
    font-size: 1.2rem;
  }
  
  .list-enter-active,
  .list-leave-active {
    transition: all 0.4s ease;
  }
  .list-enter-from,
  .list-leave-to {
    opacity: 0;
    transform: translateX(30px);
  }
  </style>