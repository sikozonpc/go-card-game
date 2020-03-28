import { useEffect, useState } from 'react'

import { ClientData } from '../types'

/** WebSocket wrapper */
const useWebsocket = (
  messageListener: (ev: MessageEvent) => void
  ) => {
  const [ws, setWebsocket] = useState<WebSocket | null>(null)

  let timeout = 250 // Initial timeout duration as a class variable

  useEffect(() => {
    // single websocket instance for the own application and constantly trying to reconnect.
    connect()
  }, [])


  /** Establishes the connect with the websocket and also ensures constant reconnection if connection closes */
  function connect() {
    var ws = new WebSocket("ws://localhost:8082/ws")
    let connectInterval: any

    // websocket onopen event listener
    ws.onopen = () => {
      console.log("connected websocket.")

      setWebsocket(ws)

      timeout = 250 // reset timer to 250 on open of websocket connection 
      clearTimeout(connectInterval) // clear Interval on on open of websocket connection
    }

    // websocket onclose event listener
    ws.onclose = e => {
      console.log(
        `Socket is closed. Reconnect will be attempted in ${Math.min(
          10000 / 1000,
          (timeout + timeout) / 1000
        )} second.`,
        e.reason
      )

      timeout += timeout //increment retry interval
      connectInterval = setTimeout(check, Math.min(10000, timeout)) //call check function after timeout
    }

    // websocket onerror event listener
    ws.onerror = (wb: any) => {
      console.error(
        "Socket encountered error: ",
        wb,
        "Closing socket"
      )

      ws.close()
    }

    ws.onmessage = (ev: MessageEvent) => {
      messageListener(ev)
    }
  }



  function sendMessage(message: ClientData) {
    if (!ws) {
      console.warn("No web socket")
      return
    }

    const data = JSON.stringify(message)

    try {
      ws.send(data)
    } catch (error) {
      console.log(error)
    }
  }

  /**
   * utilited by the @function connect to check if the connection is close, if so attempts to reconnect
   */
  function check() {
    //check if websocket instance is closed, if so call `connect` function.
    if (!ws || ws.readyState == WebSocket.CLOSED) connect()
  }

  return {
    sendMessage,

  }
}

export default useWebsocket