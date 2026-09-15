package voyage

import "github.com/DrDejaVuNG/go-ai/pkg/provider"

func optsHeaders(opts *provider.EmbedModelOptions) map[string]string {
	if opts == nil {
		return nil
	}
	return opts.Headers
}
