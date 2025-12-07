package spotify

var searchResponse = `{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=2&query=i%27m%20invisible&type=track&market=jp&locale=en-US,en;q%3D0.5",
    "limit": 2,
    "next": "https://api.spotify.com/v1/search?offset=2&limit=2&query=i%27m%20invisible&type=track&market=jp&locale=en-US,en;q%3D0.5",
    "offset": 0,
    "previous": null,
    "total": 1,
    "items": [
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/6mEQK9m2krja6X1cfsAjfl"
              },
              "href": "https://api.spotify.com/v1/artists/6mEQK9m2krja6X1cfsAjfl",
              "id": "6mEQK9m2krja6X1cfsAjfl",
              "name": "Ado",
              "type": "artist",
              "uri": "spotify:artist:6mEQK9m2krja6X1cfsAjfl"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/7Ixqxq13tWhrbnIabk3172"
          },
          "href": "https://api.spotify.com/v1/albums/7Ixqxq13tWhrbnIabk3172",
          "id": "7Ixqxq13tWhrbnIabk3172",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b2730cbecafa929898c82adc519c"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e020cbecafa929898c82adc519c"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d000048510cbecafa929898c82adc519c"
            }
          ],
          "is_playable": true,
          "name": "UTA'S SONGS ONE PIECE FILM RED",
          "release_date": "2022-08-10",
          "release_date_precision": "day",
          "total_tracks": 8,
          "type": "album",
          "uri": "spotify:album:7Ixqxq13tWhrbnIabk3172"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/6mEQK9m2krja6X1cfsAjfl"
            },
            "href": "https://api.spotify.com/v1/artists/6mEQK9m2krja6X1cfsAjfl",
            "id": "6mEQK9m2krja6X1cfsAjfl",
            "name": "Ado",
            "type": "artist",
            "uri": "spotify:artist:6mEQK9m2krja6X1cfsAjfl"
          }
        ],
        "disc_number": 1,
        "duration_ms": 257079,
        "explicit": false,
        "external_ids": {
          "isrc": "JPPO02202252"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/7yQnZkAxSiAavC0zFRB1NI"
        },
        "href": "https://api.spotify.com/v1/tracks/7yQnZkAxSiAavC0zFRB1NI",
        "id": "7yQnZkAxSiAavC0zFRB1NI",
        "is_local": false,
        "is_playable": true,
        "name": "I’m invincible",
        "popularity": 60,
        "preview_url": null,
        "track_number": 2,
        "type": "track",
        "uri": "spotify:track:7yQnZkAxSiAavC0zFRB1NI"
      }
    ]
  }
}`
