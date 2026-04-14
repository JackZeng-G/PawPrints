import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'home', component: () => import('../views/home/HomeView.vue') },
    { path: '/pets', name: 'pet-list', component: () => import('../views/pet/PetListView.vue') },
    { path: '/pets/new', name: 'pet-new', component: () => import('../views/pet/PetFormView.vue') },
    { path: '/pets/:id', name: 'pet-detail', component: () => import('../views/pet/PetDetailView.vue') },
    { path: '/pets/:id/edit', name: 'pet-edit', component: () => import('../views/pet/PetFormView.vue') },
    { path: '/diaries', name: 'diary-list', component: () => import('../views/diary/DiaryListView.vue') },
    { path: '/diaries/new', name: 'diary-new', component: () => import('../views/diary/DiaryFormView.vue') },
    { path: '/diaries/:id', name: 'diary-detail', component: () => import('../views/diary/DiaryDetailView.vue') },
    { path: '/diaries/:id/edit', name: 'diary-edit', component: () => import('../views/diary/DiaryFormView.vue') },
    { path: '/health/:petId', name: 'health', component: () => import('../views/health/HealthView.vue') },
    { path: '/reminders', name: 'reminders', component: () => import('../views/reminder/ReminderView.vue') },
    { path: '/settings', name: 'settings', component: () => import('../views/settings/SettingsView.vue') },
  ],
})

export default router
