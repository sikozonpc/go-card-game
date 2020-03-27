import React, { useState, useEffect } from 'react'
import useWebsocket from '../../hooks/useWebsocketHook'

const Battle: React.FC = () => {
  const { sendMessage } = useWebsocket(messageListener)

  const [sessionData, setSessionData] = useState(null)

  function messageListener(ev: MessageEvent) {
    setSessionData(JSON.parse(ev.data))
  }

  useEffect(() => {
    console.log('got new sessionData: ', sessionData)
  }, [sessionData])

  return (
    <div>
      <button onClick={() => sendMessage({ 
        action: '@NEW-GAME',
        data: {
          players: [
            {
              name: 'Tiago',
            },
            {
              name: 'Bot 1',
            }
          ]
        }
      })}>Hellp</button>
    </div>
  )
}

export default Battle