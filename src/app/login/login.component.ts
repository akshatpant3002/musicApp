// login.component.ts
import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable, of, switchMap } from 'rxjs';
import { SpotifyService } from 'src/spotifyservice';
import { from } from 'rxjs';
interface Playlist {
  id: string;
  name: string;
  songs: Song[];
}

interface Group {
  groupID: string;
  users: string[];
}

interface User {
  Friends:    string[];       
	LikedSong: string[];     
	UserID:     string;         
  groupAdmin: Record<string, boolean>;
  
}

interface Song {
  id: string;
  name: string;
  imageUrl: string,
  audioUrl: string
}

const url = 'http://localhost:8000/groupPost';

const groupData = {
  "groupID": 'exampleGroup',
  "users": ['user1', 'user2', 'user3']
};


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
  email: string = '';
  password: string = '';
  userData: any;
  selectedPlaylistId: string = '';
  blendPID: string = '';
  Fplaylists: Array<Playlist> = [];
  playlists: Array<Playlist> = [];
  songs: Array<Song> = [];
  userIds: string[] = [];
  user: string = '';
  dropDown: string = ''; 
  friendsUserId: string[] =[];
  selectedFriend: string = '';
  selectedPlaylistF = '';
  semiMatchedSongs: Array<Song> = []; // make a get request to pull all of the semi matched songs to display to the user
  matchedOwner: Array<Song> = [];// makes get request and allow user to add matched songs to the playlist
  randomP: Playlist[] = [];
  randomActualSongs: Playlist | undefined;
  myNumber: number = 0;
  playlistS: string = ''; // universal friend playlist that was selected
  myPlaylist: Playlist | undefined;
  arbAddSongs: Array<Song> = [];


  displaySong() { //displays a new song
    const playlistContainer = document.querySelector('.song-card');
    if (!playlistContainer) {
      console.error('Error: playlistContainer not found');
      return;
    }
  
    if (this.randomActualSongs?.songs) {
      const song = this.randomActualSongs.songs[this.myNumber];
      console.log('Displaying song:', song.name);
  
      const songContainer = document.createElement('div');
      songContainer.classList.add('song-container');
  
      const songImage = document.createElement('img');
      songImage.src = song.imageUrl;

      const audioPlayer = document.createElement('audio');
      audioPlayer.controls = true;
      audioPlayer.src = song.audioUrl;
      audioPlayer.style.position = 'absolute'; 
      audioPlayer.style.left = '820px';
      audioPlayer.style.top = '280px';
  
      const buttonContainer = document.createElement('div');
      buttonContainer.style.position = 'relative'; // Position the button container relative to the song container
      buttonContainer.style.display = 'flex';
      buttonContainer.style.justifyContent = 'space-between'; // Position the buttons at the ends of the button container
      buttonContainer.style.width = '100%'; // Make the button container the same width as the song image
  
      const likeButton = document.createElement('button');
      likeButton.innerText = 'Like';
      likeButton.style.position = 'absolute'; // Position the like button absolute to the button container
      likeButton.style.bottom = '0'; // Position the like button at the bottom of the button container
      likeButton.style.right = '0'; // Position the like button at the right of the button container
      likeButton.addEventListener('click', () => {

        this.arbAddSongs[0] = song;
        const songsA = {
          uris: this.arbAddSongs.map(song => `spotify:track:${song.id}`)
        };
        
        
        if(this.myPlaylist){
          this.spotifyService.addSongs(this.user,this.myPlaylist.id,songsA).subscribe(response => {
            });

        }

        if(this.randomActualSongs){
          // Increment myNumber to get the next song
          this.myNumber = (this.myNumber + 1) % this.randomActualSongs.songs.length;
          // Remove the current song container
          songContainer.remove();
          // Display the next song
          this.displaySong();
        }
      });
  
      const dislikeButton = document.createElement('button');
      dislikeButton.innerText = 'Dislike';
      dislikeButton.style.position = 'absolute'; // Position the dislike button absolute to the button container
      dislikeButton.style.bottom = '0'; // Position the dislike button at the bottom of the button container
      dislikeButton.style.left = '0'; // Position the dislike button at the left of the button container
      dislikeButton.addEventListener('click', () => {
        if(this.randomActualSongs){
          
          // Increment myNumber to get the next song
          this.myNumber = (this.myNumber + 1) % this.randomActualSongs.songs.length;
          // Remove the current song container
          songContainer.remove();
          // Display the next song
          this.displaySong();
        }
      });
  
      buttonContainer.appendChild(likeButton);
      buttonContainer.appendChild(dislikeButton);
  
      songContainer.appendChild(audioPlayer);
      songContainer.appendChild(songImage);
      songContainer.appendChild(buttonContainer);
  
      playlistContainer.appendChild(songContainer);
    } else {
      console.error('Error: No songs found in this.randomActualSongs');
    }
  }
  
  
  
  
  

  




  
  constructor(private router: Router, private http: HttpClient,private spotifyService: SpotifyService){
  }


  ngOnInit() {
    

  }
  addFriend(name: string) {
    if (name) {
      this.friendsUserId.push(name);
    }
  }



  onPlaylistSelected(playlistS: string) { // when friend playlist is selected new playlist is created and friend playlist is saved
    this.playlistS = playlistS;
  }



  onFriendSelected(friend: string) { // selects a friend
      console.log('Friend selected:', friend);
      this.spotifyService.getPlaylists(friend).subscribe(playlists => {
      this.Fplaylists = playlists;
      this.selectedFriend = friend;

      

    });

    
    
  }




async getUserData() { // blend fucnitonality - currently just adds both selected playlists to 

  
    this.spotifyService.createPlaylist('vibeShareTest2023',this.userIds, this.user).then(() => {

    this.spotifyService.getPlaylists(this.user).subscribe(playlists => {
    this.playlists = playlists;
    const selectedPlaylist = this.playlists.find(playlist => playlist.name === 'vibeShareTest2023');
    this.myPlaylist = selectedPlaylist;
    console.log(selectedPlaylist?.name); // does not work here
    const blendP =  this.Fplaylists.find(playlist => playlist.name === this.playlistS);
    if (blendP) {     
      //console.log(blendP.name);
    }
    
    if (selectedPlaylist && blendP) {
        
        console.log(blendP.name); // works
        this.selectedPlaylistId = selectedPlaylist.id;
        //console.log(this.selectedPlaylistId);
        //this.blendPID = blendP.id;
        this.spotifyService.getSongs(this.selectedFriend, blendP.id).subscribe(allSongs => {
        //const playlista: Playlist = this.Fplaylists.find(p => p.id === this.blendPID) as Playlist; // NOT PROPERLY MAPPING SONGS BUT FINDS IT
        //console.log(playlista.name); // FAILS HERE IN FRIEND SIDE
        console.log('WORKED');

      if (blendP && allSongs) { // DOES NOT WORK HERE
         console.log('PLAYLIST A IS READDD');
          const songs = {
            uris: allSongs.map(song => `spotify:track:${song.id}`)
          };
          this.spotifyService.addSongs(this.selectedFriend, this.selectedPlaylistId, songs) // ERROR HERE
          .subscribe(response => {
      });
      }
      });
      }
    });

    console.log('end of friends mesh')

    
    console.log(this.user)
    this.spotifyService.getPlaylists(this.user).subscribe(playlists => {
      this.playlists = playlists;
      console.log(this.playlists);
      const selectedPlaylist = this.playlists.find(playlist => playlist.name === 'vibeShareTest2023');
      this.myPlaylist = selectedPlaylist;
      console.log(this.dropDown);
      const blendP =  this.playlists.find(playlist => playlist.name === this.dropDown);
      console.log(selectedPlaylist?.id) // not detecting playlist throwbacks
      console.log(blendP?.id)
      //console.log(selectedPlaylist?.id)
      if (selectedPlaylist && blendP) {
          this.selectedPlaylistId = selectedPlaylist.id;
          this.blendPID = blendP.id;
          this.spotifyService.getSongs(this.user, this.blendPID).subscribe(allSongs => {
          const playlist = this.playlists.find(p => p.id === this.selectedPlaylistId);
          console.log(playlist?.name)

        if (playlist && allSongs) {
            console.log(allSongs);
            const songs = {
              uris: allSongs.map(song => `spotify:track:${song.id}`)
            };
            this.spotifyService.addSongs(this.user, this.selectedPlaylistId, songs) // ERROR HERE
            .subscribe(response => {
            console.log(response);
        });
        }
        });
        }
    });

  })
}

handleAuth() { // gets acess token

  //ashah03122003@gmail.com
  //ashah03122003@gmail.com
  this.spotifyService.handleAuthorizationResponse().then(() => {
    console.log(this.email)
    console.log(this.password)

        this.spotifyService.getUserId()
          .subscribe(user => {
            this.user = user;
            this.spotifyService.getPlaylists(this.user).subscribe(playlists => {
              this.playlists = playlists;
            });
          });
    });
  
}


redirectToSpotifyAPI(){
  this.spotifyService.authorize();

}


addPost(): void {
  const body = {
    groupID: "AKSHAWTYY",
    users: this.userIds
  };
  const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
  const options = { headers: headers };

  this.http.post<Group>('http://localhost:8000/groupPost', body, options)
    .subscribe(
        (res) => console.log(res),
        (err) => console.log(err)
    );
}


// type User struct {
// 	Friends    []string        `json:"friends"`
// 	LikedSong  []string        `json:"likedSong"`
// 	UserID     string          `json:"userID"`
// 	GroupAdmin map[string]bool `json:"groupAdmin"` // key is group id, admin status
// }

addUser(): void {
  console.log('RAN POST');

  const body = {
    Friends: this.friendsUserId,
    //LikedSong: ["song1", "song2", "song3"],
    userID: this.user,
    //groupAdmin: { "BINGO": true, "GROUP2": false }
  };
  const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
  const options = { headers: headers };

  this.http.post<User>('http://localhost:8000/userPost', body, options)
    .subscribe(
        (res) => console.log(res),
        (err) => console.log(err)
    );
}


  onSubmit() { // displays the songs and makes post request

    this.addPost();
    this.addUser();


    this.spotifyService.getRandomSongsFromRapCategory().subscribe(playlists => { // PRETTY CONFIDENT ABT THIS BUT NOT ANYTHING ELSE
      this.randomP = playlists;
    });

    const firstPlaylist = this.randomP[0];
    console.log(firstPlaylist); //does not work during first try - works second try
    console.log(firstPlaylist.songs);
    

    const playlistContainer = document.querySelector('.song-card')
    this.randomActualSongs = this.randomP[0]; // something wrong here

    console.log(this.randomActualSongs.songs); //does not work at all 

    //ERROR HERE:  this.randomP does not update immdietly for use afterwards. this.randomActual songs also does not update immdietly for use afterwards. please fix this code

    if (!playlistContainer) {
      console.error('No .playlist-container element was found in the DOM.');
      return;
    }
    

    
    
    if (this.randomActualSongs?.songs) {
      const song = this.randomActualSongs.songs[this.myNumber];
      if (song) {
        console.log('RAN ONCE');
        const songContainer = document.createElement('div');
        songContainer.classList.add('song-container');
        
        const songImage = document.createElement('img');
        songImage.src = song.imageUrl;
        
        const audioPlayer = document.createElement('audio');
        audioPlayer.controls = true;
        audioPlayer.src = song.audioUrl;
        audioPlayer.style.position = 'absolute'; 
        audioPlayer.style.left = '820px';
        audioPlayer.style.top = '280px';
        
    
        const buttonsContainer = document.createElement('div');
        buttonsContainer.style.position = 'relative';
        buttonsContainer.style.display = 'flex';
        buttonsContainer.style.flexDirection = 'column';
        buttonsContainer.classList.add('buttons-container');
    
        const likeButton = document.createElement('button');
        likeButton.innerText = 'Like';
        likeButton.style.position = 'absolute';
        likeButton.style.bottom = '0';
        likeButton.style.right = '0';
        likeButton.addEventListener('click', () => {
          this.arbAddSongs[0] = song;
          const songsA = {
            uris: this.arbAddSongs.map(song => `spotify:track:${song.id}`)
          };
          if(this.myPlaylist){
              this.spotifyService.addSongs(this.user,this.myPlaylist.id,songsA).subscribe(response => {
          });

          }
          if (this.randomActualSongs?.songs) {
            // Increment myNumber to get the next song
            this.myNumber = (this.myNumber + 1) % this.randomActualSongs?.songs.length;
            // Remove the current song container
            songContainer.remove();
            // Display the next song
            this.displaySong();
          }
        });
    
        const dislikeButton = document.createElement('button');
        dislikeButton.innerText = 'Dislike';
        dislikeButton.style.position = 'absolute';
        dislikeButton.style.bottom = '0';
        dislikeButton.style.left = '0';
        dislikeButton.addEventListener('click', () => {
          if (this.randomActualSongs?.songs) {
            // Increment myNumber to get the next song
            this.myNumber = (this.myNumber + 1) % this.randomActualSongs.songs.length;
            // Remove the current song container
            songContainer.remove();
            // Display the next song
            this.displaySong();
          }
        });
    
        buttonsContainer.appendChild(likeButton);
        buttonsContainer.appendChild(dislikeButton);
    
        songContainer.appendChild(audioPlayer);
        songContainer.appendChild(songImage);
        //songContainer.appendChild(audioPlayer);
        songContainer.appendChild(buttonsContainer);
    
        playlistContainer.appendChild(songContainer);
      }
    } else {
      console.error('this.randomActualSongs.songs is undefined');
    }
    
    
    }



  navigateToNewScreen() {
    this.router.navigate(['/home-screen']);
  }
}