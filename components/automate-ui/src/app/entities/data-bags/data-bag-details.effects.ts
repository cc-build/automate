import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of as observableOf } from 'rxjs';
import { catchError, mergeMap, map } from 'rxjs/operators';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';
import { DataBagItems } from './data-bags.model';
import {
  GetDataBagItems,
  GetDataBagItemsSuccess,
  GetDataBagItemsFailure,
  DataBagItemsActionTypes,
  UpdateDataBagItem,
  UpdateDataBagItemFailure,
  UpdateDataBagItemSuccess,
  DataBagItemsSuccessPayload
} from './data-bag-details.actions';

import { DataBagsRequests } from './data-bags.requests';

@Injectable()
export class DataBagItemsEffects {
  constructor(
    private actions$: Actions,
    private requests: DataBagsRequests
  ) { }

  @Effect()
  getDataBagItems$ = this.actions$.pipe(
    ofType(DataBagItemsActionTypes.GET_ALL),
    mergeMap(( action: GetDataBagItems) =>
      this.requests.getDataBagItems(action.payload).pipe(
        map((resp: DataBagItemsSuccessPayload) => new GetDataBagItemsSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new GetDataBagItemsFailure(error))))));

  @Effect()
  getDataBagItemsFailure$ = this.actions$.pipe(
    ofType(DataBagItemsActionTypes.GET_ALL_FAILURE),
    map(({ payload }: GetDataBagItemsFailure) => {
      const msg = payload.error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get infra data bag items: ${msg || payload.error}`
      });
    }));

  @Effect()
  updateDataBagItem$ = this.actions$.pipe(
    ofType(DataBagItemsActionTypes.UPDATE),
    mergeMap(({ payload: { dataBagItem } }: UpdateDataBagItem) =>
      this.requests.updateDataBagItem(dataBagItem).pipe(
        map((resp: DataBagItems) => new UpdateDataBagItemSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new UpdateDataBagItemFailure(error))))));

  @Effect()
  updateDataBagItemSuccess$ = this.actions$.pipe(
    ofType(DataBagItemsActionTypes.UPDATE_SUCCESS),
    map(({ payload: dataBagItem }: UpdateDataBagItemSuccess) => new CreateNotification({
    type: Type.info,
    message: `Successfully Updated Data Bag Item - ${dataBagItem.name}.`
  })));

  @Effect()
  updateDataBagItemFailure$ = this.actions$.pipe(
    ofType(DataBagItemsActionTypes.UPDATE_FAILURE),
    map(({ payload }: UpdateDataBagItemFailure) => {
      const msg = payload.error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not update data bag item: ${msg || payload.error}`
      });
    }));
}
