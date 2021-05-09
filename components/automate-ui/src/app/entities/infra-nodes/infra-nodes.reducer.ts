import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import { set, pipe } from 'lodash/fp';

import { EntityStatus } from 'app/entities/entities';
import { NodeActionTypes, NodeActions } from './infra-nodes.actions';
import { InfraNode } from './infra-nodes.model';

export interface InfraNodeEntityState extends EntityState<InfraNode> {
  nodesStatus: EntityStatus;
  getAllStatus: EntityStatus;
  getStatus: EntityStatus.notLoaded
  nodeList: {
    items: InfraNode[],
    total: number
  };
}

const GET_ALL_STATUS = 'getAllStatus';
const GET_STATUS = 'getStatus';

export const nodeEntityAdapter: EntityAdapter<InfraNode> = createEntityAdapter<InfraNode>({
  selectId: (infraNode: InfraNode) => infraNode.name
});

export const InfraNodeEntityInitialState: InfraNodeEntityState =
  nodeEntityAdapter.getInitialState(<InfraNodeEntityState>{
    getAllStatus: EntityStatus.notLoaded
  });

export function infraNodeEntityReducer(
  state: InfraNodeEntityState = InfraNodeEntityInitialState,
  action: NodeActions): InfraNodeEntityState {

  switch (action.type) {
    case NodeActionTypes.GET_ALL:
      return set(GET_ALL_STATUS, EntityStatus.loading, nodeEntityAdapter.removeAll(state));

    case NodeActionTypes.GET_ALL_SUCCESS:
      return pipe(
        set(GET_ALL_STATUS, EntityStatus.loadingSuccess),
        set('nodeList.items', action.payload.nodes || []),
        set('nodeList.total', action.payload.total || 0)
        )(state) as InfraNodeEntityState;

    case NodeActionTypes.GET_ALL_FAILURE:
      return set(GET_ALL_STATUS, EntityStatus.loadingFailure, state);

    case NodeActionTypes.GET:
      return set(
        GET_STATUS,
        EntityStatus.loading,
        nodeEntityAdapter.removeAll(state)
      ) as InfraNodeEntityState;

    case NodeActionTypes.GET_SUCCESS:
      return set(GET_STATUS, EntityStatus.loadingSuccess,
        nodeEntityAdapter.addOne(action.payload, state));

    case NodeActionTypes.GET_FAILURE:
      return set(GET_STATUS, EntityStatus.loadingFailure, state);

    default:
      return state;
  }
}

export const getEntityById = (id: string) =>
  (state: InfraNodeEntityState) => state.entities[id];
