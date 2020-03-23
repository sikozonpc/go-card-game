import React from 'react'
import './WebsocketWrapper.css'

interface AppState {
  ws: WebSocket | null,
}

class App extends React.Component<{}, AppState> {
  constructor(props: any) {
      super(props)

      this.state = {
          ws: null
      }
  }

  // single websocket instance for the own application and constantly trying to reconnect.

  componentDidMount() {
      this.connect()
  }

  timeout = 250 // Initial timeout duration as a class variable

  /**
   * @function connect
   * This function establishes the connect with the websocket and also ensures constant reconnection if connection closes
   */
  connect = () => {
      var ws = new WebSocket("ws://localhost:8082/ws")
      let that = this // cache the this
      let connectInterval: any

      // websocket onopen event listener
      ws.onopen = () => {
          console.log("connected websocket main component")

          this.setState({ ws: ws })

          that.timeout = 250 // reset timer to 250 on open of websocket connection 
          clearTimeout(connectInterval) // clear Interval on on open of websocket connection
      }

      // websocket onclose event listener
      ws.onclose = e => {
          console.log(
              `Socket is closed. Reconnect will be attempted in ${Math.min(
                  10000 / 1000,
                  (that.timeout + that.timeout) / 1000
              )} second.`,
              e.reason
          )

          that.timeout = that.timeout + that.timeout //increment retry interval
          connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)) //call check function after timeout
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
  }

  sendMessage = () => {
    const websocket = this.state.ws

    if (!websocket) {
      console.warn("No web socket")
      return
    }
    
    const data = JSON.stringify({ "id": "12", "name": "bidoof", "damage": 12, "health": 50 })

    try {
        websocket.send(data) //send data to the server
    } catch (error) {
        console.log(error) // catch error
    }
}

  /**
   * utilited by the @function connect to check if the connection is close, if so attempts to reconnect
   */
  check = () => {
      const { ws } = this.state
      if (!ws || ws.readyState == WebSocket.CLOSED) this.connect() //check if websocket instance is closed, if so call `connect` function.
  }

  render() {
      return <div>
        Hello world

        <button onClick={this.sendMessage}>Send message</button>
      </div>
  }
}

export default App