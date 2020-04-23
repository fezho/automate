import { map, switchMap } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { Actions, Effect, ofType } from '@ngrx/effects';
// import { NgrxStateAtom } from 'app/ngrx.reducers';
// import { Store } from '@ngrx/store';
// import { of } from 'rxjs';
// import { CreateNotification } from 'app/entities/notifications/notification.actions';
// import { Type } from 'app/entities/notifications/notification.model';

import * as actions from './nodes.actions';
// import { NodesRequests } from './nodes.requests';

import { HttpClient } from '@angular/common/http';
// import { Nodes } from './nodes.model';

import { environment } from '../../../environments/environment';

@Injectable()
export class NodesEffects {
  constructor(
    private actions$: Actions,
    // private requests: NodesRequests,
    // private store$: Store<NgrxStateAtom>,
    private httpClient: HttpClient
  ) { }


  @Effect()
  listNodes$ = this.actions$.pipe(
    ofType(actions.LIST_NODES),
    switchMap((_action) => {

      const url = `${environment.nodes_url}/search`;
      return this.httpClient.post(url, '{}');
    }),
    map(actions.listNodesSuccess));


  // @Effect()
  // listNodesFailure$ = this.actions$.pipe(
  //   ofType(actions.LIST_NODES_FAILURE),
  //   map(({ payload: { error } }: ListNodesFailure) => {
  //     const msg = error.error;
  //     return new CreateNotification({
  //       type: Type.error,
  //       message: `Could not get nodes errors: ${msg || error}`
  //     });
}