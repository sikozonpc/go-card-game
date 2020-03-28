export type ClientData = {
  Data: any,
  action: '@NEW-GAME' | '@MOVE-CARD-TO-BATTLEFIELD',
}

export type CardProps = {
  data: CardType,
}

export type CardType ={
  health: string,
  damage: string,
  id: string,
  name: string,
}