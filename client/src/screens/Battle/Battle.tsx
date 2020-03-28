import React, { useState, useEffect } from 'react'
import useWebsocket from '../../hooks/useWebsocketHook'
import Card from '../../components/Card'
import { CardType } from '../../../types'

const Battle: React.FC = () => {
  const { sendMessage } = useWebsocket(messageListener)

  const [sessionData, setSessionData] = useState<any>()

  function messageListener(ev: MessageEvent) {
    setSessionData(JSON.parse(ev.data))
  }

  useEffect(() => {
    console.log('got new sessionData: ', sessionData)
  }, [sessionData])

  const startNewGame = () => {
    sendMessage({ 
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
    })
  }

  const playerOne = sessionData && sessionData.Data.PlayerOne
  const playerTwo = sessionData && sessionData.Data.PlayerTwo

  if (!playerOne || !playerTwo) return <button onClick={startNewGame}>Start new game</button>

  return (
    <div className='Battlefield'>
      <div className='Hand'>
        {playerOne.Hand.map((card: CardType) => <Card data={card} key={`${card.id}-playerCard`} />)}
      </div>

      <div className='Hand'>
        {playerTwo.Hand.map((card: CardType) => <Card data={card} key={`${card.id}-enemyCard`} />)}
      </div>
      
    </div>
  )
}

export default Battle