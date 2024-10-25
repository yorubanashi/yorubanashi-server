// Code generated by "go run cmd/gen/main.go". Do not manually edit.

package server

import "context"

func (s *Server) cnSongsHandler(ctx context.Context, decode func(interface{}) error) (interface{}, error) {
	in := &SongRequest{}
	err := decode(in)
	if err != nil {
		return nil, err
	}
	return s.cnSongs(ctx, in)
}

func (s *Server) cnArtistsHandler(ctx context.Context, decode func(interface{}) error) (interface{}, error) {
	in := &ArtistRequest{}
	err := decode(in)
	if err != nil {
		return nil, err
	}
	return s.cnArtists(ctx, in)
}

func (s *Server) svelteWalkHandler(ctx context.Context, decode func(interface{}) error) (interface{}, error) {
	in := &SvelteWalkRequest{}
	err := decode(in)
	if err != nil {
		return nil, err
	}
	return s.svelteWalk(ctx, in)
}
