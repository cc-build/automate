import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import { set, pipe } from 'lodash/fp';
import { EntityStatus } from 'app/entities/entities';
import { ClientActionTypes, ClientActions } from './client.action';
import { Client } from './client.model';
import { HttpErrorResponse } from '@angular/common/http';

export interface ClientEntityState extends EntityState<Client> {
  clientsStatus: EntityStatus;
  getAllStatus: EntityStatus;
  clientList: {
    items: Client[],
    total: number
  };
  saveStatus: EntityStatus;
  saveError: HttpErrorResponse;
  createClient: {
    client_key: Object,
    name: string
  }
}

const GET_ALL_STATUS = 'getAllStatus';
const SAVE_STATUS = 'saveStatus';
const SAVE_ERROR = 'saveError';

export const clientEntityAdapter: EntityAdapter<Client> =
  createEntityAdapter<Client>({
  selectId: (client: Client) => client.name
});

export const ClientEntityInitialState: ClientEntityState =
  clientEntityAdapter.getInitialState(<ClientEntityState>{
  getAllStatus: EntityStatus.notLoaded
});

export function clientEntityReducer(
  state: ClientEntityState = ClientEntityInitialState,
  action: ClientActions): ClientEntityState {

  switch (action.type) {
    case ClientActionTypes.GET_ALL:
      return set(GET_ALL_STATUS, EntityStatus.loading, clientEntityAdapter.removeAll(state));

    case ClientActionTypes.GET_ALL_SUCCESS:
      return pipe(
        set('clientList.items', action.payload.clients || []),
        set('clientList.total', action.payload.total || 0)
        )(state) as ClientEntityState;

    case ClientActionTypes.GET_ALL_FAILURE:
      return set(GET_ALL_STATUS, EntityStatus.loadingFailure, state);

      case ClientActionTypes.CREATE: {
        return set(SAVE_STATUS, EntityStatus.loading, state) as ClientEntityState;
      }
  
      case ClientActionTypes.CREATE_SUCCESS: {
        return pipe(
          set('createClient.client_key', action.payload.client_key || []),
          set('createClient.name', action.payload.name || '')
        )(state) as ClientEntityState;
        // return set(SAVE_STATUS, EntityStatus.loadingSuccess, state) as ClientEntityState;
      }
  
      case ClientActionTypes.CREATE_FAILURE: {
        return pipe(
          set(SAVE_ERROR, action.payload),
          set(SAVE_STATUS, EntityStatus.loadingFailure)
        )(state) as ClientEntityState;
      }
      
    default:
      return state;
  }
}

export const getEntityById = (id: string) =>
  (state: ClientEntityState) => state.entities[id];
