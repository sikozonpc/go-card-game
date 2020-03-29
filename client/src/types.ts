export type ClientData = {
  Data: any,
  Action: '@NEW-GAME' | '@MOVE-CARD-TO-BATTLEFIELD',
}

export enum EntityTypes {
  CARD = 'CARD',
  BATTLEFIELD = 'BATTLEFIELD',
}

export type CardProps = {
  data: CardType,
  
  onDropEvent?: () => void,
}

export type CardType = {
  health: string,
  damage: string,
  id: string,
  name: string,
}

export interface BattlefieldProps {
  cards: {
    playerOne: CardType[],
    playerTwo: CardType[],
  }
}

export interface PlayerLobbyProps {
  players: any[],

  onGameStart: () => void,
}