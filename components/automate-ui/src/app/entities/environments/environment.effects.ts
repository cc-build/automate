import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of as observableOf } from 'rxjs';
import { catchError, mergeMap, map, filter } from 'rxjs/operators';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { HttpStatus } from 'app/types/types';
import { Type } from 'app/entities/notifications/notification.model';

import {
  CreateEnvironment,
  CreateEnvironmentSuccess,
  CreateEnvironmentFailure,
  GetEnvironments,
  GetEnvironmentsSuccess,
  GetEnvironmentsSuccessPayload,
  GetEnvironmentsFailure,
  EnvironmentActionTypes,
  GetEnvironment,
  GetEnvironmentSuccess,
  GetEnvironmentFailure
} from './environment.action';

import { EnvironmentRequests, EnvironmentResponse } from './environment.requests';

@Injectable()
export class EnvironmentEffects {
  constructor(
    private actions$: Actions,
    private requests: EnvironmentRequests
  ) { }

  @Effect()
  getEnvironments$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.GET_ALL),
    mergeMap((action: GetEnvironments) =>
      this.requests.getEnvironments(action.payload).pipe(
        map((resp: GetEnvironmentsSuccessPayload) => new GetEnvironmentsSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new GetEnvironmentsFailure(error))))));

  @Effect()
  getEnvironmentsFailure$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.GET_ALL_FAILURE),
    map(({ payload }: GetEnvironmentsFailure) => {
      const msg = payload.error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get infra Environment details: ${msg || payload.error}`
      });
    }));

  @Effect()
  getEnvironment$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.GET),
    mergeMap(({ payload: { server_id, org_id, name } }: GetEnvironment) =>
      this.requests.getEnvironment(server_id, org_id, name).pipe(
        map((resp) => new GetEnvironmentSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
        observableOf(new GetEnvironmentFailure(error))))));

  @Effect()
  getEnvironmentFailure$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.GET_FAILURE),
    map(({ payload }: GetEnvironmentFailure) => {
      const msg = payload.error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get environment: ${msg || payload.error}`
      });
    }));

  @Effect()
  createEnvironment$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.CREATE),
    mergeMap(({ payload: { server_id, org_id, environment } }: CreateEnvironment) =>
      this.requests.createEnvironment(server_id, org_id, environment).pipe(
        map((resp: EnvironmentResponse) => new CreateEnvironmentSuccess(resp)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new CreateEnvironmentFailure(error))))));

  @Effect()
  createEnvironmentSuccess$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.CREATE_SUCCESS),
    map(({ payload: { } }: CreateEnvironmentSuccess) => new CreateNotification({
      type: Type.info,
      message: 'Created environment.'
    })));

  @Effect()
  createEnvironmentFailure$ = this.actions$.pipe(
    ofType(EnvironmentActionTypes.CREATE_FAILURE),
    filter(({ payload }: CreateEnvironmentFailure) => payload.status !== HttpStatus.CONFLICT),
    map(({ payload }: CreateEnvironmentFailure) => new CreateNotification({
      type: Type.error,
      message: `Could not create notification: ${payload.error.error || payload}.`
    })));
}
