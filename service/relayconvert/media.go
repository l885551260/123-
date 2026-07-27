package relayconvert

import relaymedia "github.com/Project Contributors/new-api/service/relayconvert/internal/media"

type MediaResolver = relaymedia.MediaResolver

func SetMediaResolver(resolver MediaResolver) {
	relaymedia.SetMediaResolver(resolver)
}
