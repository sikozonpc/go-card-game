import React from 'react'
import { useDrop } from 'react-dnd'
import { EntityTypes, BattlefieldProps } from '../../types'
import Card from '../Card'

const Battlefield: React.FC<BattlefieldProps> = ({ cards, ...rest }) => {
  const [{ canDrop, isOver }, drop] = useDrop({
    accept: EntityTypes.CARD,
    drop: () => ({ name: EntityTypes.BATTLEFIELD }),
    collect: (monitor) => ({
      isOver: monitor.isOver(),
      canDrop: monitor.canDrop(),
    }),
  })

  const { playerOne, playerTwo } = cards

  if (!playerOne || !playerTwo) return <p>Loading...</p>


  const isActive = canDrop && isOver
  let backgroundColor = '#222'
  if (isActive) {
    backgroundColor = 'darkgreen'
  } else if (canDrop) {
    backgroundColor = 'darkkhaki'
  }

  return (
    <div className='Battlefield' ref={drop} style={{ backgroundColor }} {...rest}>
      {isActive ? 'Release to drop' : 'Drag a box here'}

      <h2>Player One</h2>
      <div className='Battlefield-area'>
        {playerOne.map(card => <Card data={card} />)}
      </div>
      <hr />
      <h2>Player Two</h2>
      <div className='Battlefield-area'>
        {playerTwo.map(card => <Card data={card} />)}
      </div>

    </div>
  )
}

export default Battlefield