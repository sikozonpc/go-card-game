import React, { useState, useEffect } from 'react'
import { DndProvider } from 'react-dnd'
import Backend from 'react-dnd-html5-backend'

import useWebsocket from '../../hooks/useWebsocketHook'
import Card from '../../components/Card'
import { CardType } from '../../types'
import Battlefield from '../../components/Battlefield'

const Battle: React.FC = () => {
  const [sessionData, setSessionData] = useState<any>()

  const messageListener = (ev: MessageEvent) => {
    setSessionData(JSON.parse(ev.data))
  }

  const { sendMessage } = useWebsocket(messageListener)


  useEffect(() => {
    console.log('got new sessionData: ', sessionData)
  }, [sessionData])

  const startNewGame = () => {
    sendMessage({
      Action: '@NEW-GAME',
      Data: {
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

  //TODO:
  const startAttack = () => { }

  const onCardBattlefieldDropHandler = (card: CardType) => () => {
    sendMessage({
      Action: '@MOVE-CARD-TO-BATTLEFIELD', Data: {
        Card: card,
        To: 'PlayerOne',
      }
    })
  }

  const playerOne = sessionData && sessionData.Data.PlayerOne
  const playerTwo = sessionData && sessionData.Data.PlayerTwo

  if (!playerOne || !playerTwo) return <button onClick={startNewGame}>Start new game</button>

  return (
    <DndProvider backend={Backend}>
      <div className='Game'>
        <button onClick={startAttack}>Attack</button>

        <div className='Hand'>
          {playerOne.Hand.map((c: CardType, idx: number) =>
            <Card data={c} key={`${idx}-${c.id}-playerCard`} onDropEvent={onCardBattlefieldDropHandler(c)} />)}
        </div>

        <Battlefield cards={{
          playerOne: playerOne.Battlefield,
          playerTwo: playerTwo.Battlefield,
        }} />

        <div className='Hand'>
          {playerTwo.Hand.map((c: CardType, idx: number) =>
            <Card data={c} key={`${idx}-${c.id}-enemyCard`} onDropEvent={onCardBattlefieldDropHandler(c)} />)}
        </div>
      </div>
    </DndProvider>
  )
}

export default Battle