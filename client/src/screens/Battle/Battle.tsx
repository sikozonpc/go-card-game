import React, { useState, useEffect } from 'react'
import { DndProvider } from 'react-dnd'
import Backend from 'react-dnd-html5-backend'

import useWebsocket from '../../hooks/useWebsocketHook'
import Card from '../../components/Card'
import { CardType, PlayerLobbyProps } from '../../types'
import Battlefield from '../../components/Battlefield'

const Battle: React.FC<{gameData: any }> = ({ gameData }) => {
  const [sessionData, setSessionData] = useState<any>(gameData)

  const messageListener = (ev: MessageEvent) => {
    setSessionData(JSON.parse(ev.data))
  }

  const { sendMessage } = useWebsocket("ws://localhost:8083/ws" , messageListener)


  useEffect(() => {
    console.log('got new sessionData: ', sessionData)
  }, [sessionData])

  //TODO:
  const joinBattleHandler = () => {}

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

  if (sessionData === true) return <p>load</p>

  const playerOne = sessionData && sessionData.Data.PlayerOne
  const playerTwo = sessionData && sessionData.Data.PlayerTwo

  if (!playerOne || !playerTwo) return <PlayerLobby onGameStart={joinBattleHandler} players={[]} />

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

const PlayerLobby: React.FC<PlayerLobbyProps> = ({ onGameStart, players }) => {
  return (
    <div>
      <button onClick={onGameStart}>START BATTLE</button>

      <h2>Players lists</h2>

      <ul>
        {players.map(p => <p>{p}</p>)}
      </ul>
    </div>
  )
}

export default Battle