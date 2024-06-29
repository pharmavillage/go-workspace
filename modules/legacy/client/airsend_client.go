package client

import (
    i00e7ba29ac45dccf3fbe13bcba3cabae2833e628561c739974d699b7c1ede884 "kiota_airsend/client/systeminfo"
    i0130c5fa219178f867adff8de62a884e3686b10bbb970983d87e7d123e2b95d4 "kiota_airsend/client/channelusersetrole"
    i020c3bbbd57d333744f166476552d96471ada9c3e2a13bef8243a2e75a0482e9 "kiota_airsend/client/userprofileset"
    i03481da376d582915ca3219bf3f7437e38880aa958fec2850567ac498e5d182e "kiota_airsend/client/channelunfavorite"
    i03f7457ca151d9cbc5f8b62f7e32863030e78f7cb57ac9df4a417be0f1d9e497 "kiota_airsend/client/wopiedit"
    i054b76c9a904d19c5e40bd167ce0bb96d5c5aae16965a145f12a22baa1ff2b8a "kiota_airsend/client/filelist"
    i09aff4fb76b09b90a4bb408f64e16131a0fa9822b6b1a4da0228d399f07757af "kiota_airsend/client/contact_formdelete"
    i0cfee8f63430377a6c0855cde640e6979326af9fce860fb4b586a21329e1fb12 "kiota_airsend/client/systemping"
    i0f53987aa6bdc355afd0ae79f5c12b112f68bd8b81157ee8eaf6271491e9dd78 "kiota_airsend/client/channelgroupdelete"
    i0ffcdd917adef76aeff5299f64b78f2399b2e846420605b665abb4aceaa5371a "kiota_airsend/client/callstatus"
    i12c6bc6df126d3baa7f05b337bafd53652c0a9874e2ab8393e8bf9334e549757 "kiota_airsend/client/lockrelease"
    i190bc560543d779b718a934be9bdc29b588519efa767ef28026fc6ca3ff3aed9 "kiota_airsend/client/usercreate"
    i1baf94b01bd32317f2d20a7f64bf12af34d78df9ba39c0f3164e9cf322a26f40 "kiota_airsend/client/actionupdate"
    i1eb5f186e7a08326244bb10d301438f990f30602f0db5fee956827dd17edb43d "kiota_airsend/client/contact_formupdate"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i277463459b579bbe4dec0d365de35c869cdefc36ae43b4922cf5a5f57d45c2b0 "kiota_airsend/client/callend"
    i27bb794c5c1b9d796641adaaced1d4d1672defdd7adf4da3623a8b829ac535bb "kiota_airsend/client/actionmove"
    i27ea2b8c34141a1d55871870fb6a126c94269a93441e531f5d344088f3d3e09d "kiota_airsend/client/fileupload"
    i29b7f5dc0ad93a5b0accfb177aa0d1957548b9eac63b412f3378f8b5d0d67736 "kiota_airsend/client/channellist"
    i29df616e16774d8ecb0b64c8b16c1f3134a2394afb69e4103bd7d231b632a6f1 "kiota_airsend/client/channelgroupadd"
    i29e4e9228696805a660755de45fd350e1c21d9d3725b5a621664f0c7a3f763d3 "kiota_airsend/client/rtmconnect"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i2c39674883c655cfe1c6fdd654954d9a7df1ffd5df02b1adf3696f202c311189 "kiota_airsend/client/channelupdate"
    i2c73fc5f98dccf0d8840640b29a723f71771e56e253c7ce33e5e726fb6a1a698 "kiota_airsend/client/userlogin"
    i2c79a7ec5171c36fc3bcdeaf94ea239de2c5d3b34ccf40773afd30f2856034aa "kiota_airsend/client/chatreactmessage"
    i2d4c12b98cb7643ab231559d7a3fc1097b43ef432f990cf66cb5fae1b1aebee7 "kiota_airsend/client/filezip"
    i30586dee8bf2008edc982a22e551688a97e70c4dd765dedaf483d026477c0be0 "kiota_airsend/client/channelactivate"
    i327d7ae381209fed6e0c8e679ea88a4b39804882948133c4675e85d2f0520a66 "kiota_airsend/client/passwordupdate"
    i34a966a01cb7ce04f704d68246e40a162c0c40d485dda5989dd7968d928041dd "kiota_airsend/client/channelremove"
    i3551a2f09cfc90bab3eab765b1e946a9559eb5c33dadfd84e059a8a2e0b27d7c "kiota_airsend/client/firebaseconnect"
    i357013209de6d20062b646a7bfb8068ec967dea478da8f3e29bd52ca514a5dd4 "kiota_airsend/client/callcreate"
    i367c8118609014d9e7be2f94c5fadf1f34a87e1e136a8a794ee3b5a4d87b4b5f "kiota_airsend/client/channelrename"
    i3864d2b2ad62f58682c507e508168fc3491f4cc200d45b61b76aaed8d4ddc090 "kiota_airsend/client/passwordreset"
    i3a07a13afcdd33a68565799fb5137eb696ad6d7088cad63dc6316773c901cbab "kiota_airsend/client/useralertack"
    i3e1cdd17cce0571a000c15284567b6c7cc461ff8a74af5816f0ee26282cfca68 "kiota_airsend/client/oauthserverclientcreate"
    i461e20c9fdf4d56f2bc9ab2580a75a3046b0af8e3004c7f4d951e3c3862d28b6 "kiota_airsend/client/userverify"
    i474597017ec1b483a5911d83cd7c7d93804a9385efc09fb89696b2922063d6d4 "kiota_airsend/client/filecopy"
    i485bec29c9ea34457ac549d9f5b905dda2ced48c30c53cf23f2efb90e26db271 "kiota_airsend/client/wikipreview"
    i48b01e15b14502e41b5b093ef4650315bb42f2b727046c9419f98d72fe3ad762 "kiota_airsend/client/fileversions"
    i4b94ba300b92289caefcd679b8dd307b4e64e9fb9b17c738af479c79221b24dd "kiota_airsend/client/lockrefresh"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i4ed31338e9dcb0e2852c4fcf63a845c3c0d17d9917f2fb54301e80f1e15dbb8c "kiota_airsend/client/oauthserverapprove"
    i534b4e695afbeab2ddd5d33b40e6bc1e31ba17320ee226dc8c951d9d778162bb "kiota_airsend/client/fileinfo"
    i5367062788e7ef3bd5e7c6b8938174fa6059490f2810d1ad555828e3a12dc50f "kiota_airsend/client/contact_formfill"
    i53edbe5f3c4a76ac5a43ed1e1abbcfc6671c9f08eefe229e8aada0be499fc72c "kiota_airsend/client/oauthserverclientlist"
    i550c07614fec91c001c652c0877a803e81691cddf9ce5e3c3891d51d7c6368c7 "kiota_airsend/client/lockacquire"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i5796d786143e0d9d32cd80dff0e8c79512af3329fe400e4796bbcdaac22aad39 "kiota_airsend/client/filedownload"
    i5bce865ae640694b7fce7d4cc421de5b7eb220e3124fde71fa6f4318e9f939cb "kiota_airsend/client/firebasedisconnect"
    i61b72ec4e08ef5f3b636cb1ebed2beb14396f22770bcef39ee477234168b3124 "kiota_airsend/client/actiondelete"
    i6312452da4936cec9806d78ff45eeca8ec086971d16064afa8a07b9cae8e634b "kiota_airsend/client/userfinalize"
    i63615b8b772f09618a8a9f0fa91a3f5480cc3b7f359e43d9ec6364469917b827 "kiota_airsend/client/channelgroupmove"
    i644802ad2e34e5b84b9c28db328f193529545af76c20ee03b72c6465ba661990 "kiota_airsend/client/channelimageset"
    i653c94699810e3a61ec72c31f14d2f895166a960f79e478667e9f602e2a657d0 "kiota_airsend/client/filecreate"
    i689be25abcd60e50fc7d37c8615aea575ae4cce827b50716928290a53a336f22 "kiota_airsend/client/oauthserverauthorize"
    i68f8b7e06b4829d5cf81b80505d548e41378f21cad5901fdf1d9464db100c860 "kiota_airsend/client/usernotificationsreport"
    i69d99e6f46b4a50100ac519ce48ab16030117bd644f27c1a4b133043cccf8fed "kiota_airsend/client/channelgroupremove"
    i6aed5909142d0b27cfb1ce8cb4c6990272a44e9c77cf93c3234934436432e6de "kiota_airsend/client/channeljoin"
    i6c24363e68bf22fa37f772ed03da2c0a0010b0512f3969a5f97eba358a30bf8a "kiota_airsend/client/userlogout"
    i6c4bc6612cfc9797b11ffe2450d6cbf871f40aed14af1196e7f309545297aaa0 "kiota_airsend/client/filedelete"
    i6f2ac62185eaca2fedd843d04b622b9c4f9cce22c5d10c96b959b5394a2cff61 "kiota_airsend/client/actionlist"
    i720955961a2deffb7979c64d56d2338fbe2fc87b6f37fcdc3da3309ce660aff8 "kiota_airsend/client/actioncreate"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i730cd8e18456e4864b57a44fa776043f898f8d3e07da6f1e3fb6ea88f05996eb "kiota_airsend/client/channelcreate"
    i75a5336f31c03d6bc5b0b7d391e03db0e43ed72d16c1fd9d448270e377c53e00 "kiota_airsend/client/wopiview"
    i78a2a5c3e8e48a8fe68e88c55af6a6eb857fde1108c3f5b346436e831b0d4bee "kiota_airsend/client/userimageget"
    i7a619dc264d9859c28458df373531da015a10cf858aa49d6d130220cc584ab9f "kiota_airsend/client/channelexport"
    i7bd4d78dee014ac3f03d72cfccf9db15fefa9a7345da827c9ee9c5e5115d4770 "kiota_airsend/client/channelclose"
    i7c1f1d5a22a2b12646080149c775b81bca5c7d2a22d836e5885002f6ebbc5d34 "kiota_airsend/client/channelfavorite"
    i7dbf718a85b0a42a3d7888826968e2e3018e5dc1aa12d5948c83093e0465c88b "kiota_airsend/client/chatcommand"
    i7eca5803237bb046ab47f9fa855fbbbad81f233a4cbc0b67f4615b31066507ec "kiota_airsend/client/channelgrouplist"
    i7eff6a5806a956858764b4aa4dd8101a3af096c991a41d3f2e5834d0183f9b71 "kiota_airsend/client/callinvite"
    i88a1cff3f2b380cbaf94d161692292d9e549dedbb1e16454f24641be30be15b7 "kiota_airsend/client/oauthserveraccess_token"
    i8a00381ceab459a841534a4cf720f36316d68553759e64f6c9122c8177d87286 "kiota_airsend/client/callinviteaccept"
    i8b62ba227b97b418596a7f51e7a31ffe069a5cf6291a524f2f37791bb8649609 "kiota_airsend/client/channelkick"
    i8d9fdc760d4f08db5274b0ebf9166885f12159a07a406bbc699c69f6dd956ca3 "kiota_airsend/client/useralerts"
    i962f17b331592aade147c527cf8762477e99d7eeafb65a0e2d27fa9c86246b95 "kiota_airsend/client/channelinfo"
    i9818475d286f327d978eff2754e00a8d669d072cd03afb1a36dac1606c08cc1d "kiota_airsend/client/channelimageget"
    i983981b8478f71a01703599304ecfb2dde36433e611489391458122fffbba59c "kiota_airsend/client/contact_formcreate"
    i98abdf60faf242d9836d007982544357773c6d800fa7f7970a712c2f9d4141f5 "kiota_airsend/client/channelreadnotification"
    i9e0332588261bb677203e6540a738ae982f296e31b99887dcad170e128147f5d "kiota_airsend/client/passwordrecover"
    i9e99e72ad904d3d5ba06c79812c5231ed05279550751cfe4a0b9ec043e5bf99a "kiota_airsend/client/channelinvite"
    ia188c31289c0df2d9cef9f3acd914c62ff09ee505bc84e182a371a622791b586 "kiota_airsend/client/channelmembers"
    ia2c2da8a4faf4f522f0b496fa8f87d9eef829734e24dd9032bb0e33824f5c056 "kiota_airsend/client/actioninfo"
    iad4a5e156c2d4c9048aa826ec7d30a69656248a1f987d88e4429d6931a6667c8 "kiota_airsend/client/chatpostmessage"
    iae6b29890a4f80efa6f6f516204a88944f707005d849306591235a5f16bcb912 "kiota_airsend/client/usernotificationsmanage"
    ib82dd1eee2ef40d745de3d295db3df3a30f3a8856b18bc3811a5da88ebae40ff "kiota_airsend/client/filethumb"
    ic44875f819936ee980882175efcc49eba30ab122851a25bf7535cc5c77119726 "kiota_airsend/client/callupdate"
    ic491edf41a81aa1a7446daf35f6e98d0db761de3414719b6970221c0775dc363 "kiota_airsend/client/filemove"
    ic4d8c7999e657bbef5d78946c647f2d2360f7abbbb20aa8f32a85ac34923875d "kiota_airsend/client/channelnotificationsmanage"
    ic532668acaef9bcf89528de310fb996171b98f8567b86c3671314cbc66b4cddc "kiota_airsend/client/contact_formlist"
    ic61c5d2daffdb21f33bedf4f92ed3a427667d3875b4eab45f72c2914dbe66984 "kiota_airsend/client/channeloneonone"
    ic92907e506aa781fa7b7aa302c93cbbe3b04bf07c710d65c054002d5cf08cc6e "kiota_airsend/client/userloginrefresh"
    ica7d969bd49a9681acb4d6d49c5a4a1d5ebc439d0fb1b9a48beec4734d8c6098 "kiota_airsend/client/wikiget"
    icd47069403ffb5fe37d203f645cbfd0cb3f4056f26aa8ad788d64c49de44b397 "kiota_airsend/client/channelhistory"
    id84d956212ae4d922e757c430aee7f422028eaf1ebe1a612f071caa6f0d42386 "kiota_airsend/client/searchquery"
    idfaf2a7775b1ffe8b3613612d68de7fd94b37cadff6be39282a2790c9973a1a0 "kiota_airsend/client/channelgroupupdate"
    ie0e789b44cb14b373b5fcb81df62334a26a74e704543f3028de5377d5b01d394 "kiota_airsend/client/channelgroupcreate"
    ie957ea31bc7c7521c35bff5c6e1293dff8089f66f041b8aaed00bee5cad0d5f2 "kiota_airsend/client/userimageset"
    ieb7eed8a19394bb4c045488a2d3598189a978c92ab762d167ea62af97afc49d0 "kiota_airsend/client/userinfo"
    iec59f7055fd0ca61eeeceba736b399401547734b639a90f6ce9b3fda2e776665 "kiota_airsend/client/calljoin"
    ied96d6c21cd670c678082a836b9da55998417163c7c4726929ad629ed6afb493 "kiota_airsend/client/channelleave"
    ief486f5920a250f477819bc92ef70eae572b99fc5c584977c4964e745bbf0b1d "kiota_airsend/client/chatupdatemessage"
    ifb0d3aedda1432d359a83fe5e6894ff971f7e7f8a385cbe0af82d1c40a5f7729 "kiota_airsend/client/chatdeletemessage"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AirsendClient the main entry point of the SDK, exposes the configuration and the fluent API.
type AirsendClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActionCreate the actionCreate property
func (m *AirsendClient) ActionCreate()(*i720955961a2deffb7979c64d56d2338fbe2fc87b6f37fcdc3da3309ce660aff8.ActionCreateRequestBuilder) {
    return i720955961a2deffb7979c64d56d2338fbe2fc87b6f37fcdc3da3309ce660aff8.NewActionCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ActionDelete the actionDelete property
func (m *AirsendClient) ActionDelete()(*i61b72ec4e08ef5f3b636cb1ebed2beb14396f22770bcef39ee477234168b3124.ActionDeleteRequestBuilder) {
    return i61b72ec4e08ef5f3b636cb1ebed2beb14396f22770bcef39ee477234168b3124.NewActionDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ActionInfo the actionInfo property
func (m *AirsendClient) ActionInfo()(*ia2c2da8a4faf4f522f0b496fa8f87d9eef829734e24dd9032bb0e33824f5c056.ActionInfoRequestBuilder) {
    return ia2c2da8a4faf4f522f0b496fa8f87d9eef829734e24dd9032bb0e33824f5c056.NewActionInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ActionList the actionList property
func (m *AirsendClient) ActionList()(*i6f2ac62185eaca2fedd843d04b622b9c4f9cce22c5d10c96b959b5394a2cff61.ActionListRequestBuilder) {
    return i6f2ac62185eaca2fedd843d04b622b9c4f9cce22c5d10c96b959b5394a2cff61.NewActionListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ActionMove the actionMove property
func (m *AirsendClient) ActionMove()(*i27bb794c5c1b9d796641adaaced1d4d1672defdd7adf4da3623a8b829ac535bb.ActionMoveRequestBuilder) {
    return i27bb794c5c1b9d796641adaaced1d4d1672defdd7adf4da3623a8b829ac535bb.NewActionMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ActionUpdate the actionUpdate property
func (m *AirsendClient) ActionUpdate()(*i1baf94b01bd32317f2d20a7f64bf12af34d78df9ba39c0f3164e9cf322a26f40.ActionUpdateRequestBuilder) {
    return i1baf94b01bd32317f2d20a7f64bf12af34d78df9ba39c0f3164e9cf322a26f40.NewActionUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallCreate the callCreate property
func (m *AirsendClient) CallCreate()(*i357013209de6d20062b646a7bfb8068ec967dea478da8f3e29bd52ca514a5dd4.CallCreateRequestBuilder) {
    return i357013209de6d20062b646a7bfb8068ec967dea478da8f3e29bd52ca514a5dd4.NewCallCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallEnd the callEnd property
func (m *AirsendClient) CallEnd()(*i277463459b579bbe4dec0d365de35c869cdefc36ae43b4922cf5a5f57d45c2b0.CallEndRequestBuilder) {
    return i277463459b579bbe4dec0d365de35c869cdefc36ae43b4922cf5a5f57d45c2b0.NewCallEndRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallInvite the callInvite property
func (m *AirsendClient) CallInvite()(*i7eff6a5806a956858764b4aa4dd8101a3af096c991a41d3f2e5834d0183f9b71.CallInviteRequestBuilder) {
    return i7eff6a5806a956858764b4aa4dd8101a3af096c991a41d3f2e5834d0183f9b71.NewCallInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallInviteAccept the callInviteAccept property
func (m *AirsendClient) CallInviteAccept()(*i8a00381ceab459a841534a4cf720f36316d68553759e64f6c9122c8177d87286.CallInviteAcceptRequestBuilder) {
    return i8a00381ceab459a841534a4cf720f36316d68553759e64f6c9122c8177d87286.NewCallInviteAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallJoin the callJoin property
func (m *AirsendClient) CallJoin()(*iec59f7055fd0ca61eeeceba736b399401547734b639a90f6ce9b3fda2e776665.CallJoinRequestBuilder) {
    return iec59f7055fd0ca61eeeceba736b399401547734b639a90f6ce9b3fda2e776665.NewCallJoinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallStatus the callStatus property
func (m *AirsendClient) CallStatus()(*i0ffcdd917adef76aeff5299f64b78f2399b2e846420605b665abb4aceaa5371a.CallStatusRequestBuilder) {
    return i0ffcdd917adef76aeff5299f64b78f2399b2e846420605b665abb4aceaa5371a.NewCallStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CallUpdate the callUpdate property
func (m *AirsendClient) CallUpdate()(*ic44875f819936ee980882175efcc49eba30ab122851a25bf7535cc5c77119726.CallUpdateRequestBuilder) {
    return ic44875f819936ee980882175efcc49eba30ab122851a25bf7535cc5c77119726.NewCallUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelActivate the channelActivate property
func (m *AirsendClient) ChannelActivate()(*i30586dee8bf2008edc982a22e551688a97e70c4dd765dedaf483d026477c0be0.ChannelActivateRequestBuilder) {
    return i30586dee8bf2008edc982a22e551688a97e70c4dd765dedaf483d026477c0be0.NewChannelActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelClose the channelClose property
func (m *AirsendClient) ChannelClose()(*i7bd4d78dee014ac3f03d72cfccf9db15fefa9a7345da827c9ee9c5e5115d4770.ChannelCloseRequestBuilder) {
    return i7bd4d78dee014ac3f03d72cfccf9db15fefa9a7345da827c9ee9c5e5115d4770.NewChannelCloseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelCreate the channelCreate property
func (m *AirsendClient) ChannelCreate()(*i730cd8e18456e4864b57a44fa776043f898f8d3e07da6f1e3fb6ea88f05996eb.ChannelCreateRequestBuilder) {
    return i730cd8e18456e4864b57a44fa776043f898f8d3e07da6f1e3fb6ea88f05996eb.NewChannelCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelExport the channelExport property
func (m *AirsendClient) ChannelExport()(*i7a619dc264d9859c28458df373531da015a10cf858aa49d6d130220cc584ab9f.ChannelExportRequestBuilder) {
    return i7a619dc264d9859c28458df373531da015a10cf858aa49d6d130220cc584ab9f.NewChannelExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelFavorite the channelFavorite property
func (m *AirsendClient) ChannelFavorite()(*i7c1f1d5a22a2b12646080149c775b81bca5c7d2a22d836e5885002f6ebbc5d34.ChannelFavoriteRequestBuilder) {
    return i7c1f1d5a22a2b12646080149c775b81bca5c7d2a22d836e5885002f6ebbc5d34.NewChannelFavoriteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupAdd the channelGroupAdd property
func (m *AirsendClient) ChannelGroupAdd()(*i29df616e16774d8ecb0b64c8b16c1f3134a2394afb69e4103bd7d231b632a6f1.ChannelGroupAddRequestBuilder) {
    return i29df616e16774d8ecb0b64c8b16c1f3134a2394afb69e4103bd7d231b632a6f1.NewChannelGroupAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupCreate the channelGroupCreate property
func (m *AirsendClient) ChannelGroupCreate()(*ie0e789b44cb14b373b5fcb81df62334a26a74e704543f3028de5377d5b01d394.ChannelGroupCreateRequestBuilder) {
    return ie0e789b44cb14b373b5fcb81df62334a26a74e704543f3028de5377d5b01d394.NewChannelGroupCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupDelete the channelGroupDelete property
func (m *AirsendClient) ChannelGroupDelete()(*i0f53987aa6bdc355afd0ae79f5c12b112f68bd8b81157ee8eaf6271491e9dd78.ChannelGroupDeleteRequestBuilder) {
    return i0f53987aa6bdc355afd0ae79f5c12b112f68bd8b81157ee8eaf6271491e9dd78.NewChannelGroupDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupList the channelGroupList property
func (m *AirsendClient) ChannelGroupList()(*i7eca5803237bb046ab47f9fa855fbbbad81f233a4cbc0b67f4615b31066507ec.ChannelGroupListRequestBuilder) {
    return i7eca5803237bb046ab47f9fa855fbbbad81f233a4cbc0b67f4615b31066507ec.NewChannelGroupListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupMove the channelGroupMove property
func (m *AirsendClient) ChannelGroupMove()(*i63615b8b772f09618a8a9f0fa91a3f5480cc3b7f359e43d9ec6364469917b827.ChannelGroupMoveRequestBuilder) {
    return i63615b8b772f09618a8a9f0fa91a3f5480cc3b7f359e43d9ec6364469917b827.NewChannelGroupMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupRemove the channelGroupRemove property
func (m *AirsendClient) ChannelGroupRemove()(*i69d99e6f46b4a50100ac519ce48ab16030117bd644f27c1a4b133043cccf8fed.ChannelGroupRemoveRequestBuilder) {
    return i69d99e6f46b4a50100ac519ce48ab16030117bd644f27c1a4b133043cccf8fed.NewChannelGroupRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelGroupUpdate the channelGroupUpdate property
func (m *AirsendClient) ChannelGroupUpdate()(*idfaf2a7775b1ffe8b3613612d68de7fd94b37cadff6be39282a2790c9973a1a0.ChannelGroupUpdateRequestBuilder) {
    return idfaf2a7775b1ffe8b3613612d68de7fd94b37cadff6be39282a2790c9973a1a0.NewChannelGroupUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelHistory the channelHistory property
func (m *AirsendClient) ChannelHistory()(*icd47069403ffb5fe37d203f645cbfd0cb3f4056f26aa8ad788d64c49de44b397.ChannelHistoryRequestBuilder) {
    return icd47069403ffb5fe37d203f645cbfd0cb3f4056f26aa8ad788d64c49de44b397.NewChannelHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelImageGet the channelImageGet property
func (m *AirsendClient) ChannelImageGet()(*i9818475d286f327d978eff2754e00a8d669d072cd03afb1a36dac1606c08cc1d.ChannelImageGetRequestBuilder) {
    return i9818475d286f327d978eff2754e00a8d669d072cd03afb1a36dac1606c08cc1d.NewChannelImageGetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelImageSet the channelImageSet property
func (m *AirsendClient) ChannelImageSet()(*i644802ad2e34e5b84b9c28db328f193529545af76c20ee03b72c6465ba661990.ChannelImageSetRequestBuilder) {
    return i644802ad2e34e5b84b9c28db328f193529545af76c20ee03b72c6465ba661990.NewChannelImageSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelInfo the channelInfo property
func (m *AirsendClient) ChannelInfo()(*i962f17b331592aade147c527cf8762477e99d7eeafb65a0e2d27fa9c86246b95.ChannelInfoRequestBuilder) {
    return i962f17b331592aade147c527cf8762477e99d7eeafb65a0e2d27fa9c86246b95.NewChannelInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelInvite the channelInvite property
func (m *AirsendClient) ChannelInvite()(*i9e99e72ad904d3d5ba06c79812c5231ed05279550751cfe4a0b9ec043e5bf99a.ChannelInviteRequestBuilder) {
    return i9e99e72ad904d3d5ba06c79812c5231ed05279550751cfe4a0b9ec043e5bf99a.NewChannelInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelJoin the channelJoin property
func (m *AirsendClient) ChannelJoin()(*i6aed5909142d0b27cfb1ce8cb4c6990272a44e9c77cf93c3234934436432e6de.ChannelJoinRequestBuilder) {
    return i6aed5909142d0b27cfb1ce8cb4c6990272a44e9c77cf93c3234934436432e6de.NewChannelJoinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelKick the channelKick property
func (m *AirsendClient) ChannelKick()(*i8b62ba227b97b418596a7f51e7a31ffe069a5cf6291a524f2f37791bb8649609.ChannelKickRequestBuilder) {
    return i8b62ba227b97b418596a7f51e7a31ffe069a5cf6291a524f2f37791bb8649609.NewChannelKickRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelLeave the channelLeave property
func (m *AirsendClient) ChannelLeave()(*ied96d6c21cd670c678082a836b9da55998417163c7c4726929ad629ed6afb493.ChannelLeaveRequestBuilder) {
    return ied96d6c21cd670c678082a836b9da55998417163c7c4726929ad629ed6afb493.NewChannelLeaveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelList the channelList property
func (m *AirsendClient) ChannelList()(*i29b7f5dc0ad93a5b0accfb177aa0d1957548b9eac63b412f3378f8b5d0d67736.ChannelListRequestBuilder) {
    return i29b7f5dc0ad93a5b0accfb177aa0d1957548b9eac63b412f3378f8b5d0d67736.NewChannelListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelMembers the channelMembers property
func (m *AirsendClient) ChannelMembers()(*ia188c31289c0df2d9cef9f3acd914c62ff09ee505bc84e182a371a622791b586.ChannelMembersRequestBuilder) {
    return ia188c31289c0df2d9cef9f3acd914c62ff09ee505bc84e182a371a622791b586.NewChannelMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelNotificationsManage the channelNotificationsManage property
func (m *AirsendClient) ChannelNotificationsManage()(*ic4d8c7999e657bbef5d78946c647f2d2360f7abbbb20aa8f32a85ac34923875d.ChannelNotificationsManageRequestBuilder) {
    return ic4d8c7999e657bbef5d78946c647f2d2360f7abbbb20aa8f32a85ac34923875d.NewChannelNotificationsManageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelOneOnOne the channelOneOnOne property
func (m *AirsendClient) ChannelOneOnOne()(*ic61c5d2daffdb21f33bedf4f92ed3a427667d3875b4eab45f72c2914dbe66984.ChannelOneOnOneRequestBuilder) {
    return ic61c5d2daffdb21f33bedf4f92ed3a427667d3875b4eab45f72c2914dbe66984.NewChannelOneOnOneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelReadNotification the channelReadNotification property
func (m *AirsendClient) ChannelReadNotification()(*i98abdf60faf242d9836d007982544357773c6d800fa7f7970a712c2f9d4141f5.ChannelReadNotificationRequestBuilder) {
    return i98abdf60faf242d9836d007982544357773c6d800fa7f7970a712c2f9d4141f5.NewChannelReadNotificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelRemove the channelRemove property
func (m *AirsendClient) ChannelRemove()(*i34a966a01cb7ce04f704d68246e40a162c0c40d485dda5989dd7968d928041dd.ChannelRemoveRequestBuilder) {
    return i34a966a01cb7ce04f704d68246e40a162c0c40d485dda5989dd7968d928041dd.NewChannelRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelRename the channelRename property
func (m *AirsendClient) ChannelRename()(*i367c8118609014d9e7be2f94c5fadf1f34a87e1e136a8a794ee3b5a4d87b4b5f.ChannelRenameRequestBuilder) {
    return i367c8118609014d9e7be2f94c5fadf1f34a87e1e136a8a794ee3b5a4d87b4b5f.NewChannelRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelUnfavorite the channelUnfavorite property
func (m *AirsendClient) ChannelUnfavorite()(*i03481da376d582915ca3219bf3f7437e38880aa958fec2850567ac498e5d182e.ChannelUnfavoriteRequestBuilder) {
    return i03481da376d582915ca3219bf3f7437e38880aa958fec2850567ac498e5d182e.NewChannelUnfavoriteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelUpdate the channelUpdate property
func (m *AirsendClient) ChannelUpdate()(*i2c39674883c655cfe1c6fdd654954d9a7df1ffd5df02b1adf3696f202c311189.ChannelUpdateRequestBuilder) {
    return i2c39674883c655cfe1c6fdd654954d9a7df1ffd5df02b1adf3696f202c311189.NewChannelUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChannelUserSetrole the channelUserSetrole property
func (m *AirsendClient) ChannelUserSetrole()(*i0130c5fa219178f867adff8de62a884e3686b10bbb970983d87e7d123e2b95d4.ChannelUserSetroleRequestBuilder) {
    return i0130c5fa219178f867adff8de62a884e3686b10bbb970983d87e7d123e2b95d4.NewChannelUserSetroleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChatCommand the chatCommand property
func (m *AirsendClient) ChatCommand()(*i7dbf718a85b0a42a3d7888826968e2e3018e5dc1aa12d5948c83093e0465c88b.ChatCommandRequestBuilder) {
    return i7dbf718a85b0a42a3d7888826968e2e3018e5dc1aa12d5948c83093e0465c88b.NewChatCommandRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChatDeletemessage the chatDeletemessage property
func (m *AirsendClient) ChatDeletemessage()(*ifb0d3aedda1432d359a83fe5e6894ff971f7e7f8a385cbe0af82d1c40a5f7729.ChatDeletemessageRequestBuilder) {
    return ifb0d3aedda1432d359a83fe5e6894ff971f7e7f8a385cbe0af82d1c40a5f7729.NewChatDeletemessageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChatPostmessage the chatPostmessage property
func (m *AirsendClient) ChatPostmessage()(*iad4a5e156c2d4c9048aa826ec7d30a69656248a1f987d88e4429d6931a6667c8.ChatPostmessageRequestBuilder) {
    return iad4a5e156c2d4c9048aa826ec7d30a69656248a1f987d88e4429d6931a6667c8.NewChatPostmessageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChatReactmessage the chatReactmessage property
func (m *AirsendClient) ChatReactmessage()(*i2c79a7ec5171c36fc3bcdeaf94ea239de2c5d3b34ccf40773afd30f2856034aa.ChatReactmessageRequestBuilder) {
    return i2c79a7ec5171c36fc3bcdeaf94ea239de2c5d3b34ccf40773afd30f2856034aa.NewChatReactmessageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChatUpdatemessage the chatUpdatemessage property
func (m *AirsendClient) ChatUpdatemessage()(*ief486f5920a250f477819bc92ef70eae572b99fc5c584977c4964e745bbf0b1d.ChatUpdatemessageRequestBuilder) {
    return ief486f5920a250f477819bc92ef70eae572b99fc5c584977c4964e745bbf0b1d.NewChatUpdatemessageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAirsendClient instantiates a new AirsendClient and sets the default values.
func NewAirsendClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AirsendClient) {
    m := &AirsendClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://live.airsend.io/api/v1")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Contact_formCreate the contact_formCreate property
func (m *AirsendClient) Contact_formCreate()(*i983981b8478f71a01703599304ecfb2dde36433e611489391458122fffbba59c.Contact_formCreateRequestBuilder) {
    return i983981b8478f71a01703599304ecfb2dde36433e611489391458122fffbba59c.NewContact_formCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contact_formDelete the contact_formDelete property
func (m *AirsendClient) Contact_formDelete()(*i09aff4fb76b09b90a4bb408f64e16131a0fa9822b6b1a4da0228d399f07757af.Contact_formDeleteRequestBuilder) {
    return i09aff4fb76b09b90a4bb408f64e16131a0fa9822b6b1a4da0228d399f07757af.NewContact_formDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contact_formFill the contact_formFill property
func (m *AirsendClient) Contact_formFill()(*i5367062788e7ef3bd5e7c6b8938174fa6059490f2810d1ad555828e3a12dc50f.Contact_formFillRequestBuilder) {
    return i5367062788e7ef3bd5e7c6b8938174fa6059490f2810d1ad555828e3a12dc50f.NewContact_formFillRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contact_formList the contact_formList property
func (m *AirsendClient) Contact_formList()(*ic532668acaef9bcf89528de310fb996171b98f8567b86c3671314cbc66b4cddc.Contact_formListRequestBuilder) {
    return ic532668acaef9bcf89528de310fb996171b98f8567b86c3671314cbc66b4cddc.NewContact_formListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contact_formUpdate the contact_formUpdate property
func (m *AirsendClient) Contact_formUpdate()(*i1eb5f186e7a08326244bb10d301438f990f30602f0db5fee956827dd17edb43d.Contact_formUpdateRequestBuilder) {
    return i1eb5f186e7a08326244bb10d301438f990f30602f0db5fee956827dd17edb43d.NewContact_formUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileCopy the fileCopy property
func (m *AirsendClient) FileCopy()(*i474597017ec1b483a5911d83cd7c7d93804a9385efc09fb89696b2922063d6d4.FileCopyRequestBuilder) {
    return i474597017ec1b483a5911d83cd7c7d93804a9385efc09fb89696b2922063d6d4.NewFileCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileCreate the fileCreate property
func (m *AirsendClient) FileCreate()(*i653c94699810e3a61ec72c31f14d2f895166a960f79e478667e9f602e2a657d0.FileCreateRequestBuilder) {
    return i653c94699810e3a61ec72c31f14d2f895166a960f79e478667e9f602e2a657d0.NewFileCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileDelete the fileDelete property
func (m *AirsendClient) FileDelete()(*i6c4bc6612cfc9797b11ffe2450d6cbf871f40aed14af1196e7f309545297aaa0.FileDeleteRequestBuilder) {
    return i6c4bc6612cfc9797b11ffe2450d6cbf871f40aed14af1196e7f309545297aaa0.NewFileDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileDownload the fileDownload property
func (m *AirsendClient) FileDownload()(*i5796d786143e0d9d32cd80dff0e8c79512af3329fe400e4796bbcdaac22aad39.FileDownloadRequestBuilder) {
    return i5796d786143e0d9d32cd80dff0e8c79512af3329fe400e4796bbcdaac22aad39.NewFileDownloadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileInfo the fileInfo property
func (m *AirsendClient) FileInfo()(*i534b4e695afbeab2ddd5d33b40e6bc1e31ba17320ee226dc8c951d9d778162bb.FileInfoRequestBuilder) {
    return i534b4e695afbeab2ddd5d33b40e6bc1e31ba17320ee226dc8c951d9d778162bb.NewFileInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileList the fileList property
func (m *AirsendClient) FileList()(*i054b76c9a904d19c5e40bd167ce0bb96d5c5aae16965a145f12a22baa1ff2b8a.FileListRequestBuilder) {
    return i054b76c9a904d19c5e40bd167ce0bb96d5c5aae16965a145f12a22baa1ff2b8a.NewFileListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileMove the fileMove property
func (m *AirsendClient) FileMove()(*ic491edf41a81aa1a7446daf35f6e98d0db761de3414719b6970221c0775dc363.FileMoveRequestBuilder) {
    return ic491edf41a81aa1a7446daf35f6e98d0db761de3414719b6970221c0775dc363.NewFileMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileThumb the fileThumb property
func (m *AirsendClient) FileThumb()(*ib82dd1eee2ef40d745de3d295db3df3a30f3a8856b18bc3811a5da88ebae40ff.FileThumbRequestBuilder) {
    return ib82dd1eee2ef40d745de3d295db3df3a30f3a8856b18bc3811a5da88ebae40ff.NewFileThumbRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileUpload the fileUpload property
func (m *AirsendClient) FileUpload()(*i27ea2b8c34141a1d55871870fb6a126c94269a93441e531f5d344088f3d3e09d.FileUploadRequestBuilder) {
    return i27ea2b8c34141a1d55871870fb6a126c94269a93441e531f5d344088f3d3e09d.NewFileUploadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileVersions the fileVersions property
func (m *AirsendClient) FileVersions()(*i48b01e15b14502e41b5b093ef4650315bb42f2b727046c9419f98d72fe3ad762.FileVersionsRequestBuilder) {
    return i48b01e15b14502e41b5b093ef4650315bb42f2b727046c9419f98d72fe3ad762.NewFileVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileZip the fileZip property
func (m *AirsendClient) FileZip()(*i2d4c12b98cb7643ab231559d7a3fc1097b43ef432f990cf66cb5fae1b1aebee7.FileZipRequestBuilder) {
    return i2d4c12b98cb7643ab231559d7a3fc1097b43ef432f990cf66cb5fae1b1aebee7.NewFileZipRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FirebaseConnect the firebaseConnect property
func (m *AirsendClient) FirebaseConnect()(*i3551a2f09cfc90bab3eab765b1e946a9559eb5c33dadfd84e059a8a2e0b27d7c.FirebaseConnectRequestBuilder) {
    return i3551a2f09cfc90bab3eab765b1e946a9559eb5c33dadfd84e059a8a2e0b27d7c.NewFirebaseConnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FirebaseDisconnect the firebaseDisconnect property
func (m *AirsendClient) FirebaseDisconnect()(*i5bce865ae640694b7fce7d4cc421de5b7eb220e3124fde71fa6f4318e9f939cb.FirebaseDisconnectRequestBuilder) {
    return i5bce865ae640694b7fce7d4cc421de5b7eb220e3124fde71fa6f4318e9f939cb.NewFirebaseDisconnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LockAcquire the lockAcquire property
func (m *AirsendClient) LockAcquire()(*i550c07614fec91c001c652c0877a803e81691cddf9ce5e3c3891d51d7c6368c7.LockAcquireRequestBuilder) {
    return i550c07614fec91c001c652c0877a803e81691cddf9ce5e3c3891d51d7c6368c7.NewLockAcquireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LockRefresh the lockRefresh property
func (m *AirsendClient) LockRefresh()(*i4b94ba300b92289caefcd679b8dd307b4e64e9fb9b17c738af479c79221b24dd.LockRefreshRequestBuilder) {
    return i4b94ba300b92289caefcd679b8dd307b4e64e9fb9b17c738af479c79221b24dd.NewLockRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LockRelease the lockRelease property
func (m *AirsendClient) LockRelease()(*i12c6bc6df126d3baa7f05b337bafd53652c0a9874e2ab8393e8bf9334e549757.LockReleaseRequestBuilder) {
    return i12c6bc6df126d3baa7f05b337bafd53652c0a9874e2ab8393e8bf9334e549757.NewLockReleaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OauthServerAccess_token the oauthServerAccess_token property
func (m *AirsendClient) OauthServerAccess_token()(*i88a1cff3f2b380cbaf94d161692292d9e549dedbb1e16454f24641be30be15b7.OauthServerAccess_tokenRequestBuilder) {
    return i88a1cff3f2b380cbaf94d161692292d9e549dedbb1e16454f24641be30be15b7.NewOauthServerAccess_tokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OauthServerApprove the oauthServerApprove property
func (m *AirsendClient) OauthServerApprove()(*i4ed31338e9dcb0e2852c4fcf63a845c3c0d17d9917f2fb54301e80f1e15dbb8c.OauthServerApproveRequestBuilder) {
    return i4ed31338e9dcb0e2852c4fcf63a845c3c0d17d9917f2fb54301e80f1e15dbb8c.NewOauthServerApproveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OauthServerAuthorize the oauthServerAuthorize property
func (m *AirsendClient) OauthServerAuthorize()(*i689be25abcd60e50fc7d37c8615aea575ae4cce827b50716928290a53a336f22.OauthServerAuthorizeRequestBuilder) {
    return i689be25abcd60e50fc7d37c8615aea575ae4cce827b50716928290a53a336f22.NewOauthServerAuthorizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OauthServerClientCreate the oauthServerClientCreate property
func (m *AirsendClient) OauthServerClientCreate()(*i3e1cdd17cce0571a000c15284567b6c7cc461ff8a74af5816f0ee26282cfca68.OauthServerClientCreateRequestBuilder) {
    return i3e1cdd17cce0571a000c15284567b6c7cc461ff8a74af5816f0ee26282cfca68.NewOauthServerClientCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OauthServerClientList the oauthServerClientList property
func (m *AirsendClient) OauthServerClientList()(*i53edbe5f3c4a76ac5a43ed1e1abbcfc6671c9f08eefe229e8aada0be499fc72c.OauthServerClientListRequestBuilder) {
    return i53edbe5f3c4a76ac5a43ed1e1abbcfc6671c9f08eefe229e8aada0be499fc72c.NewOauthServerClientListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PasswordRecover the passwordRecover property
func (m *AirsendClient) PasswordRecover()(*i9e0332588261bb677203e6540a738ae982f296e31b99887dcad170e128147f5d.PasswordRecoverRequestBuilder) {
    return i9e0332588261bb677203e6540a738ae982f296e31b99887dcad170e128147f5d.NewPasswordRecoverRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PasswordReset the passwordReset property
func (m *AirsendClient) PasswordReset()(*i3864d2b2ad62f58682c507e508168fc3491f4cc200d45b61b76aaed8d4ddc090.PasswordResetRequestBuilder) {
    return i3864d2b2ad62f58682c507e508168fc3491f4cc200d45b61b76aaed8d4ddc090.NewPasswordResetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PasswordUpdate the passwordUpdate property
func (m *AirsendClient) PasswordUpdate()(*i327d7ae381209fed6e0c8e679ea88a4b39804882948133c4675e85d2f0520a66.PasswordUpdateRequestBuilder) {
    return i327d7ae381209fed6e0c8e679ea88a4b39804882948133c4675e85d2f0520a66.NewPasswordUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RtmConnect the rtmConnect property
func (m *AirsendClient) RtmConnect()(*i29e4e9228696805a660755de45fd350e1c21d9d3725b5a621664f0c7a3f763d3.RtmConnectRequestBuilder) {
    return i29e4e9228696805a660755de45fd350e1c21d9d3725b5a621664f0c7a3f763d3.NewRtmConnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchQuery the searchQuery property
func (m *AirsendClient) SearchQuery()(*id84d956212ae4d922e757c430aee7f422028eaf1ebe1a612f071caa6f0d42386.SearchQueryRequestBuilder) {
    return id84d956212ae4d922e757c430aee7f422028eaf1ebe1a612f071caa6f0d42386.NewSearchQueryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SystemInfo the systemInfo property
func (m *AirsendClient) SystemInfo()(*i00e7ba29ac45dccf3fbe13bcba3cabae2833e628561c739974d699b7c1ede884.SystemInfoRequestBuilder) {
    return i00e7ba29ac45dccf3fbe13bcba3cabae2833e628561c739974d699b7c1ede884.NewSystemInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SystemPing the systemPing property
func (m *AirsendClient) SystemPing()(*i0cfee8f63430377a6c0855cde640e6979326af9fce860fb4b586a21329e1fb12.SystemPingRequestBuilder) {
    return i0cfee8f63430377a6c0855cde640e6979326af9fce860fb4b586a21329e1fb12.NewSystemPingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserAlertAck the userAlertAck property
func (m *AirsendClient) UserAlertAck()(*i3a07a13afcdd33a68565799fb5137eb696ad6d7088cad63dc6316773c901cbab.UserAlertAckRequestBuilder) {
    return i3a07a13afcdd33a68565799fb5137eb696ad6d7088cad63dc6316773c901cbab.NewUserAlertAckRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserAlerts the userAlerts property
func (m *AirsendClient) UserAlerts()(*i8d9fdc760d4f08db5274b0ebf9166885f12159a07a406bbc699c69f6dd956ca3.UserAlertsRequestBuilder) {
    return i8d9fdc760d4f08db5274b0ebf9166885f12159a07a406bbc699c69f6dd956ca3.NewUserAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserCreate the userCreate property
func (m *AirsendClient) UserCreate()(*i190bc560543d779b718a934be9bdc29b588519efa767ef28026fc6ca3ff3aed9.UserCreateRequestBuilder) {
    return i190bc560543d779b718a934be9bdc29b588519efa767ef28026fc6ca3ff3aed9.NewUserCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserFinalize the userFinalize property
func (m *AirsendClient) UserFinalize()(*i6312452da4936cec9806d78ff45eeca8ec086971d16064afa8a07b9cae8e634b.UserFinalizeRequestBuilder) {
    return i6312452da4936cec9806d78ff45eeca8ec086971d16064afa8a07b9cae8e634b.NewUserFinalizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserImageGet the userImageGet property
func (m *AirsendClient) UserImageGet()(*i78a2a5c3e8e48a8fe68e88c55af6a6eb857fde1108c3f5b346436e831b0d4bee.UserImageGetRequestBuilder) {
    return i78a2a5c3e8e48a8fe68e88c55af6a6eb857fde1108c3f5b346436e831b0d4bee.NewUserImageGetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserImageSet the userImageSet property
func (m *AirsendClient) UserImageSet()(*ie957ea31bc7c7521c35bff5c6e1293dff8089f66f041b8aaed00bee5cad0d5f2.UserImageSetRequestBuilder) {
    return ie957ea31bc7c7521c35bff5c6e1293dff8089f66f041b8aaed00bee5cad0d5f2.NewUserImageSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserInfo the userInfo property
func (m *AirsendClient) UserInfo()(*ieb7eed8a19394bb4c045488a2d3598189a978c92ab762d167ea62af97afc49d0.UserInfoRequestBuilder) {
    return ieb7eed8a19394bb4c045488a2d3598189a978c92ab762d167ea62af97afc49d0.NewUserInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserLogin the userLogin property
func (m *AirsendClient) UserLogin()(*i2c73fc5f98dccf0d8840640b29a723f71771e56e253c7ce33e5e726fb6a1a698.UserLoginRequestBuilder) {
    return i2c73fc5f98dccf0d8840640b29a723f71771e56e253c7ce33e5e726fb6a1a698.NewUserLoginRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserLoginRefresh the userLoginRefresh property
func (m *AirsendClient) UserLoginRefresh()(*ic92907e506aa781fa7b7aa302c93cbbe3b04bf07c710d65c054002d5cf08cc6e.UserLoginRefreshRequestBuilder) {
    return ic92907e506aa781fa7b7aa302c93cbbe3b04bf07c710d65c054002d5cf08cc6e.NewUserLoginRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserLogout the userLogout property
func (m *AirsendClient) UserLogout()(*i6c24363e68bf22fa37f772ed03da2c0a0010b0512f3969a5f97eba358a30bf8a.UserLogoutRequestBuilder) {
    return i6c24363e68bf22fa37f772ed03da2c0a0010b0512f3969a5f97eba358a30bf8a.NewUserLogoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserNotificationsManage the userNotificationsManage property
func (m *AirsendClient) UserNotificationsManage()(*iae6b29890a4f80efa6f6f516204a88944f707005d849306591235a5f16bcb912.UserNotificationsManageRequestBuilder) {
    return iae6b29890a4f80efa6f6f516204a88944f707005d849306591235a5f16bcb912.NewUserNotificationsManageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserNotificationsReport the userNotificationsReport property
func (m *AirsendClient) UserNotificationsReport()(*i68f8b7e06b4829d5cf81b80505d548e41378f21cad5901fdf1d9464db100c860.UserNotificationsReportRequestBuilder) {
    return i68f8b7e06b4829d5cf81b80505d548e41378f21cad5901fdf1d9464db100c860.NewUserNotificationsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserProfileSet the userProfileSet property
func (m *AirsendClient) UserProfileSet()(*i020c3bbbd57d333744f166476552d96471ada9c3e2a13bef8243a2e75a0482e9.UserProfileSetRequestBuilder) {
    return i020c3bbbd57d333744f166476552d96471ada9c3e2a13bef8243a2e75a0482e9.NewUserProfileSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserVerify the userVerify property
func (m *AirsendClient) UserVerify()(*i461e20c9fdf4d56f2bc9ab2580a75a3046b0af8e3004c7f4d951e3c3862d28b6.UserVerifyRequestBuilder) {
    return i461e20c9fdf4d56f2bc9ab2580a75a3046b0af8e3004c7f4d951e3c3862d28b6.NewUserVerifyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WikiGet the wikiGet property
func (m *AirsendClient) WikiGet()(*ica7d969bd49a9681acb4d6d49c5a4a1d5ebc439d0fb1b9a48beec4734d8c6098.WikiGetRequestBuilder) {
    return ica7d969bd49a9681acb4d6d49c5a4a1d5ebc439d0fb1b9a48beec4734d8c6098.NewWikiGetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WikiPreview the wikiPreview property
func (m *AirsendClient) WikiPreview()(*i485bec29c9ea34457ac549d9f5b905dda2ced48c30c53cf23f2efb90e26db271.WikiPreviewRequestBuilder) {
    return i485bec29c9ea34457ac549d9f5b905dda2ced48c30c53cf23f2efb90e26db271.NewWikiPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WopiEdit the wopiEdit property
func (m *AirsendClient) WopiEdit()(*i03f7457ca151d9cbc5f8b62f7e32863030e78f7cb57ac9df4a417be0f1d9e497.WopiEditRequestBuilder) {
    return i03f7457ca151d9cbc5f8b62f7e32863030e78f7cb57ac9df4a417be0f1d9e497.NewWopiEditRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WopiView the wopiView property
func (m *AirsendClient) WopiView()(*i75a5336f31c03d6bc5b0b7d391e03db0e43ed72d16c1fd9d448270e377c53e00.WopiViewRequestBuilder) {
    return i75a5336f31c03d6bc5b0b7d391e03db0e43ed72d16c1fd9d448270e377c53e00.NewWopiViewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
