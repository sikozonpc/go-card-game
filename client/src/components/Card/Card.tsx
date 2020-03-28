import React from 'react'
import { CardProps } from '../../../types'
import Draggable, { DraggableData, DraggableEvent } from 'react-draggable'

const Card: React.FC<CardProps> = ({ data, ...rest }) => {

  function onDragStophandler(e: DraggableEvent, data: DraggableData) {
    //console.log(e, data)
  }

  return (
    <Draggable
      onStop={onDragStophandler}
      {...rest}
    >
    <div className='Card'>
      <p>{data.name}</p>

      <div className='Stats'>
      <p>HP:  {data.health}</p>
      <p>DMG: {data.damage}</p>
      </div>
    </div>
    </Draggable>
  )
}

export default Card