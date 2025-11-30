import request from './request'

// PVE处罚相关API
export const penaltyAPI = {
    // 搜索角色处罚信息
    searchPenalty(params) {
        return request({
            url: '/player/violations/info',
            method: 'GET',
            params
        })
    },

    // 获取服务器列表
    getServerList(params = {}) {
        return request({
            url: '/server/status',
            method: 'GET',
            params: {
                name: params.name,
                sort: params.sort || 'desc'
            }
        })
    },

    // 获取服务器处罚占比
    getServerPenaltyList(limit) {
        return request({
            url: `/player/violations/server-proportions?limit=${limit}`,
            method: 'GET'
        })
    },

    // 获取角色详细信息
    getCharacterDetail(characterId) {
        return request({
            url: `/character/${characterId}`,
            method: 'GET'
        })
    },

    // 获取本周词缀
    getMethicAffix() {
        return request({
            url: '/mythic/affixes',
            method: 'GET'
        })
    },

    // 获取玩家角色列表
    getPlayerList(params) {
        return request({
            url: '/search/playerList',
            method: 'GET',
            params
        })
    }
}