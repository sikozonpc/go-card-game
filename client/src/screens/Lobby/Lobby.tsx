import React, { useState } from 'react'
import { API_URL } from '../../constants'
import Battle from '../Battle/Battle'

const Lobby: React.FC = () => {

  const [game, setGame] = useState()
  const [isReady, setReady] = useState(false)

  const startNewGameHandler = () => {
    fetch(`${API_URL}/newgame`)
      .then(res => {
        return res.json()
      })
      .then(data => {
        console.log(data)

        setGame(data)
      })
      .catch(err => console.error(err))
  }

  const joinGameHandler = () => {
    fetch(`${API_URL}/game`)
    .then(res => {
      return res.json()
    })
    .then(data => {
      console.log(data)

      setGame(data)
    })
    .catch(err => console.error(err))
  }

  return (
    <div>
      <button onClick={startNewGameHandler}>New game</button>
      <button onClick={joinGameHandler}>Join current game</button>

      {game && <button onClick={() => setReady(true)}>Ready!</button>}

      {game && isReady && <Battle gameData={game} />}
    </div>
  )
}

export default Lobby