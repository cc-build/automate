import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { Store } from '@ngrx/store';
import { UserPreference } from './user-preferences.model';
import { userPreferencesTimeformat } from './user-preferences.selector';
import {
  GetUserPreferences,
  UpdateUserPreferences,
  TestUpdateUserTimeformat } from './user-preferences.actions';



@Injectable({ providedIn: 'root'})
export class UserPreferencesService {

  constructor(
    private store: Store<NgrxStateAtom>
  ) {}

  timeformat$: Observable<UserPreference> = this.store.select(userPreferencesTimeformat);

    // only for testing development - will be removed or modified before release
  testUpdateUserTimeformat(format: string) {
    this.store.dispatch(new TestUpdateUserTimeformat(format));
  }

  getUserPreferences() {
    this.store.dispatch(new GetUserPreferences());
  }

  updateUserPreferences(tz) {
    this.store.dispatch(new UpdateUserPreferences(tz));
  }

}