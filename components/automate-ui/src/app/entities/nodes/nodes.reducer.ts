import { set, pipe } from 'lodash/fp';

import * as actions from './nodes.actions';
import { Node } from './nodes.model';
import { EntityStatus } from '../entities';


export interface NodesEntityState {
  nodes: Node[];
  listNodesStatus: EntityStatus;

}

export const nodesEntityInitialState: NodesEntityState = {
  nodes: [],
  listNodesStatus: EntityStatus.notLoaded
};

export function userEntityReducer(state: NodesEntityState = nodesEntityInitialState,
  action: actions.NodesAction) {

  switch (action.type) {
    case actions.LIST_NODES:
      return set('listNodesStatus', EntityStatus.loading, state);

    case actions.LIST_NODES_SUCCESS:
      return pipe(
        set('listNodesStatus', EntityStatus.loadingSuccess),
        set('nodes', action.payload))(state);

    // case actions.LIST_NODES_FAILURE:
    //   return set('listNodesStatus', EntityStatus.loadingFailure, state);

    default:
      return state;

  }
}