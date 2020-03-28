import React, { useState } from 'react'
import { API_URL } from '../../constants'

const Lobby: React.FC = () => {

  const [game, setGame] = useState()

  const startNewGameHandler = () => {
    fetch(`${API_URL}/`)
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

      {/* {game} */}
    </div>
  )
}

export default Lobby