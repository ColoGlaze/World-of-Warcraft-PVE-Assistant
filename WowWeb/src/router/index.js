import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import SearchPenalty from '../views/SearchPenalty.vue'
import DpsSimulator from '../views/DpsSimulator.vue'
import AboutUs from '../views/AboutUs.vue'
import Profile from '../views/Profile.vue'
import ServerStatus from '../views/ServerStatus.vue'
import MdtPlanner from '../views/MdtPlanner.vue'
import CharacterSearch from '../views/CharacterSearch.vue'
import CharacterDetail from '../views/CharacterDetail.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/search',
        name: 'SearchPenalty',
        component: SearchPenalty
    },
    {
        path: '/dps-simulator',
        name: 'DpsSimulator',
        component: DpsSimulator
    },
    {
        path: '/about',
        name: 'AboutUs',
        component: AboutUs
    },
    {
        path: '/profile',
        name: 'Profile',
        component: Profile
    },
    {
        path: '/server-status',
        name: 'ServerStatus',
        component: ServerStatus
    },
    {
        path: '/mdt',
        name: 'MdtPlanner',
        component: MdtPlanner
    },
    {
        path: '/penalty-ranking',
        name: 'PenaltyRanking',
        component: () => import('../views/PenaltyRanking.vue')
    },
    {
        path: '/character-search',
        name: 'CharacterSearch',
        component: CharacterSearch
    },
    {
        path: '/character-detail',
        name: 'CharacterDetail',
        component: CharacterDetail
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
    scrollBehavior() {
        return { top: 0 }
    }
})

export default router