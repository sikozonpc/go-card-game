import React from 'react'
import { CardProps, EntityTypes } from '../../types'
import { useDrag, DragSourceMonitor } from 'react-dnd'

export type ItemProps = {
  type: string,
  name?: string
}

/** Card element */
const Card: React.FC<CardProps> = ({ data, onDropEvent, ...rest }) => {

  const [{ isDragging }, drag] = useDrag({
    item: { type: EntityTypes.CARD },
    end: (item: ItemProps | undefined, monitor: DragSourceMonitor) => {
      const dropResult = monitor.getDropResult()

      if (item && dropResult) {
        console.log(`You dropped a ${item.type} into a ${dropResult.name}!`)

        if (onDropEvent) {
          onDropEvent()
        }
      }
    },
    collect: (monitor) => ({
      isDragging: monitor.isDragging(),
    }),
  })

  const opacity = isDragging ? 0.4 : 1

  return (
    <div className='Card' ref={drag} style={{ opacity }} {...rest}>
      <p>{data.name}</p>

      <div className='Stats'>
      <p>HP:  {data.health}</p>
      <p>DMG: {data.damage}</p>
      </div>
    </div>
  )
}

export default Card