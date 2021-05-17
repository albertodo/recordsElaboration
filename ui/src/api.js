//const { vueAppApiUrl } = require('@/env')
const axios = require('axios')
 
const TIMEOUT_MS = 30000


function getFullData () {
    return new Promise(function (resolve, reject) {

      const headers = {
        'content-type': 'application/json'
      }
   
      const httpClient = axios.create({
        baseURL: "http://localhost:3000",
        headers: headers,
        timeout: TIMEOUT_MS
      })
      httpClient.get('/api/data')
        .then(function (res) {
          resolve(res.data)
        })
        .catch(function (error) {
          console.log(error)
          reject(error)
        })
    })
  }
  
  function deleteItemWithID (itemID) {
    return new Promise(function (resolve, reject) {
      console.log("Here I get: " + itemID)
      const headers = {
        'content-type': 'application/json'
      }
   
      const httpClient = axios.create({
        baseURL: "http://localhost:3000",
        headers: headers,
        timeout: TIMEOUT_MS
      })

      httpClient.delete('/api/deleteItemWithID', {data: {id: itemID}})
        .then(function (res) {
          resolve(res.data)
        })
        .catch(function (error) {
          console.log(error)
          reject(error)
        })
    })
  }

  function evaluateStats () {
    return new Promise(function (resolve, reject) {
      const headers = {
        'content-type': 'application/json'
      }
   
      const httpClient = axios.create({
        baseURL: "http://localhost:3000",
        headers: headers,
        timeout: TIMEOUT_MS
      })

      httpClient.get('/api/stats')
        .then(function (res) {
          resolve(res.data)
        })
        .catch(function (error) {
          console.log(error)
          reject(error)
        })
    })
  }
module.exports = {
  
  getFullData,
  deleteItemWithID,
  evaluateStats,
}