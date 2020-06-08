package discord

// {
// 	"username": "Discord Bot",
// 	"avatar_url": "http://google.com",
//     "content": "Welcome to <:discohook:694285995394203759>Discohook, a free message and embed builder for Discord!\nThere's additional info in the embeds below, or you can use the *Clear all* button in the editor to start making embeds.\nHave questions? Discohook has a support server at <https://discohook.org/discord>.",
//     "embeds": [
//         {
//             "title": "Legal things",
//             "description": "To make Discohook as helpful as it can be, we use some assets derived from Discord's application. Discohook has no affiliation with Discord in any way, shape, or form.\n\nThe source code to this app is [available on GitHub](https://github.com/discohook/discohook) licensed under the GNU Affero General Public License v3.0.\nIf you need to contact me, you can join the [support server](https://discohook.org/discord), or send an email to \"hello\" at discohook.org.",
//             "url": "http://dotsandbrackets.com/wp-content/uploads/2017/01/grafana-dashboard.jpg",
//             "color": 15746887,
//             "fields": [
//                 {
//                     "name": "Name field",
//                     "value": "Field value"
//                 },
//                 {
//                     "name": "Name field",
//                     "value": "Field value"
//                 }
//             ],
//             "author": {
//                 "name": "Author",
//                 "url": "http://google.com",
//                 "icon_url": "http://dotsandbrackets.com/wp-content/uploads/2017/01/grafana-dashboard.jpg"
//             },
//             "footer": {
//                 "text": "footer",
//                 "icon_url": "http://dotsandbrackets.com/wp-content/uploads/2017/01/grafana-dashboard.jpg"
//             },
//             "image": {
//                 "url": "http://dotsandbrackets.com/wp-content/uploads/2017/01/grafana-dashboard.jpg"
//             },
//             "thumbnail": {
//                 "url": "http://dotsandbrackets.com/wp-content/uploads/2017/01/grafana-dashboard.jpg"
//             }
//         }
//     ]
// }

// DiscordManOut json structure for Discord
type DiscordManOut struct {
	UserName  string             `json:"username,omitempty"`
	AvatarURL string             `json:"avatar_url,omitempty"`
	Content   string             `json:"content,omitempty"`
	Embeds    []DiscordManEmbeds `json:"embeds,omitempty"`
}

// DiscordManEmbeds json structure for Discord Message Embed
type DiscordManEmbeds struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Color       string `json:"color,omitempty"`
	Fields      []struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"fields,omitempty"`
	Author struct {
		Name    string `json:"name,omitempty"`
		URL     string `json:"url,omitempty"`
		IconURL string `json:"icon_url,omitempty"`
	} `json:"author,omitempty"`
	Footer struct {
		Text    string `json:"text,omitempty"`
		IconURL string `json:"icon_url,omitempty"`
	} `json:"footer,omitempty"`
	Image struct {
		URL string `json:"url,omitempty"`
	} `json:"image,omitempty"`
	Thumbnail struct {
		URL string `json"url,omitempty"`
	} `json:"thumbnail,omitempty"`
}
