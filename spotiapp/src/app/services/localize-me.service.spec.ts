import { TestBed } from '@angular/core/testing';

import { LocalizeMeService } from './localize-me.service';

describe('LocalizeMeService', () => {
  let service: LocalizeMeService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LocalizeMeService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
