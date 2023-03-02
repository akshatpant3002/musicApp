
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { forkJoin, Observable, of } from 'rxjs';
import { catchError, map, switchMap, tap } from 'rxjs/operators';
import { LoginComponent } from './app/login/login.component';



interface Playlist {
    songs: Song[];
    id: string;
    name: string;
    // other properties of a playlist
  }
  interface Song {
    id: string;
    name: string;
    imageUrl: string,
    audioUrl: string,

    // other properties of a song
  }
  interface SongResponse {
    id: string;
    name: string;
    imageUrl: string;
    audioUrl: string;
    // other properties of a song returned by the Spotify API
  }

  interface CategoryResponse {
    id: string;
    name: string;
    
  }



@Injectable({
  providedIn: 'root'
})
export class SpotifyService {

  accessToken: string = '';
  private apiBaseUrl = 'https://api.spotify.com/v1/';
  user: String = '';

  constructor(private http: HttpClient, private router: Router) {}
  

  authorize(): Observable<any>{
    const clientId = '3571de52a7d747358b31518e6b0e6b1f';
    const redirectUri = 'http://localhost:4200';
    const scope = 'user-read-private user-read-email playlist-modify-public playlist-modify-private playlist-read-collaborative user-library-modify';
    window.location.href = `https://accounts.spotify.com/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}&response_type=code`;
    return of({});
  }
 
  async handleAuthorizationResponse() { // gets me the access token
    
    const clientId = '3571de52a7d747358b31518e6b0e6b1f';
    const clientSecret = '1dff1fc95abd4bf28c5ef114ba7e58bb';
    const redirectUri = 'http://localhost:4200';

    const code = this.getCodeFromRedirectUri();
    if (!code) {
      return;
    }

    const headers = new HttpHeaders({
      'Authorization': `Basic ${btoa(`${clientId}:${clientSecret}`)}`,
      'Content-Type': 'application/x-www-form-urlencoded'
    });

    try {
      const response = await this.http.post<{ access_token: string}>(
        'https://accounts.spotify.com/api/token',
        `grant_type=authorization_code&code=${code}&redirect_uri=${redirectUri}`,
        { headers }
      ).toPromise();
      if (response) {
          this.accessToken = response.access_token;
      }
      this.router.navigate(['/']);
    } catch (error) {
      console.error(error);
    }
  }

  private getCodeFromRedirectUri(): string | null {
    const url = new URL(window.location.href);
    const code = url.searchParams.get('code');
    return code || null;
  }

  getUser(userId: string) {
    const headers = new HttpHeaders({
      'Authorization': `Bearer ${this.accessToken}`
    });
    return this.http.get(`https://api.spotify.com/v1/users/${userId}`, { headers });
  }


  getUserId() {
    const headers = new HttpHeaders({
      'Authorization': `Bearer ${this.accessToken}`
    });
  
    return this.http.get<any>(`https://api.spotify.com/v1/me`, { headers })
      .pipe(
        map(response => {

          return response.id;
        })
      );
  }

  getPlaylists(userId: string): Observable<Array<Playlist>> {
    const headers = new HttpHeaders({
        'Authorization': `Bearer ${this.accessToken}`
 
      });
    return this.http.get<any>(`https://api.spotify.com/v1/users/${userId}/playlists`, { headers })
      .pipe(
        map(response => {
          return response.items.map((playlist: any) => {
            return {
                songs: this.getSongs(userId, playlist.name), // IDK IF THIS IS STILL NECESSARY BUT LEAVE IT FOR NOW
                id: playlist.id,
                name: playlist.name
            } as unknown as Playlist;
          });
        })
      );
  }

  createPlaylist(name: string, userIds: string[], username: string) {

    const headers = new HttpHeaders().set('Authorization', `Bearer ${this.accessToken}`);
    console.log(username);
    const userId = username;// CHANGE THIS TO BE DYNAMIC BY SOMEHOW CALLING GETUSERID
    const url = `${this.apiBaseUrl}users/${userId}/playlists`;
    const body = {
      name: name,
      public: false,
      collaborative: true
    };

    return this.http.post<any>(url, body, { headers }).toPromise()
    //no point in adding collaborators bc only owner can edit the playlist anyways
  }


  songs = {
    uris: [
    ]
  };
  addSongs(userId: string, playlistId: string, songs: { uris: string[]; }) {
    const headers = new HttpHeaders({
      'Authorization': `Bearer ${this.accessToken}`,
      'Content-Type': 'application/json'
    });
  
    return this.http.post(`https://api.spotify.com/v1/users/${userId}/playlists/${playlistId}/tracks/`, songs, { headers });
  }

  getCategoryId(categoryName: string): Observable<string | null> {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${this.accessToken}`);

    return this.http.get<{ categories: { items: CategoryResponse[] } }>(`https://api.spotify.com/v1/browse/categories`, { headers })
  .pipe(
    map(response => {
      const categories = response.categories.items;
      const category = categories.find(categoryResponse => categoryResponse.name === categoryName);
      return category ? category.id : null;
    })
  );
  }

  getRandomSongsFromRapCategory(): Observable<Playlist[]> {
    return this.getCategoryId('Pop').pipe(
      switchMap(rapCategoryId => {
        if (!rapCategoryId) {
          return of([]);
        }
  
        const headers = new HttpHeaders({
          'Authorization': `Bearer ${this.accessToken}`
        });
  
        return this.http.get<any>(`https://api.spotify.com/v1/browse/categories/${rapCategoryId}/playlists`, { headers })
          .pipe(
            switchMap(response => {
              const playlistObservables = response.playlists.items.map((playlist: {
                uri: any; id: string; name: string; imageURL: string 
}) => {
                // Fetch the playlist ID from the 'uri' property
                const id = playlist.uri.split(':')[2];
                return this.getSongsFromPlaylist(id).pipe(
                  map(songs => ({
                    id: id,
                    name: playlist.name,
                    songs: songs
                  } as Playlist))
                );
              });
              return forkJoin(playlistObservables);
            }),
            map(playlists => playlists as Playlist[])
          );
      })
    );
  }
  
  getSongsFromPlaylist(playlistId: string): Observable<Array<Song>> {
    const headers = new HttpHeaders({
        'Authorization': `Bearer ${this.accessToken}`
      });

    return this.http.get<any>(`https://api.spotify.com/v1/playlists/${playlistId}/tracks`, { headers })
      .pipe(
        // tap((response: any) => console.log('Response:', response)),
        // catchError(error => {
        //   console.error('Error:', error);
        //   return of([]);
        // }),
        map(response => {
          return response.items.map((song: any) => {
            //console.log(song.track.name) //work on first try yayyyy
            return {
              id: song.track.id,
              name: song.track.name,
              imageUrl: song.track.album.images[0].url,
              audioUrl: song.track.preview_url
            } as Song;
          });
        })
      );
  }


  getSongs(userId: string,  playlistId: string): Observable<Array<Song>> {
    const headers = new HttpHeaders({
        'Authorization': `Bearer ${this.accessToken}`
      });
    return this.http.get<any>(`https://api.spotify.com/v1/users/${userId}/playlists/${playlistId}/tracks`, { headers })
      .pipe(
        
        map(response => {
          return response.items.map((tracks:any) => {
            return {
              id: tracks.track.id,
              name: tracks.track.name,
              imageUrl: tracks.imageUrl,
              audioUrl: tracks.audioUrl,

            } as Song;
          });
        })
        
      ); 
  }
}



