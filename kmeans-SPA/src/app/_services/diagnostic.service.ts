import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { ListDiagnostic } from '../_models/list-diagnostic';

@Injectable({
  providedIn: 'root'
})
export class DiagnosticService {


  baseUrl = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getDiagnostics(): Observable<ListDiagnostic> {
    return this.http.get<ListDiagnostic>(this.baseUrl + 'diagnosticos');
  }
}
