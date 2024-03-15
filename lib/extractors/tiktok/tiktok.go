package tiktok

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	errors2 "github.com/juju/errors"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"tutu-gin/lib/extractors"
)

func init() {
	extractors.Register("tiktok", New())
}

type extractor struct{}

type tiktokContent struct {
	StatusCode int `json:"status_code"`
	MinCursor  int `json:"min_cursor"`
	MaxCursor  int `json:"max_cursor"`
	HasMore    int `json:"has_more"`
	AwemeList  []struct {
		AwemeId    string `json:"aweme_id"`
		Desc       string `json:"desc"`
		CreateTime int    `json:"create_time"`
		Author     struct {
			Uid         string `json:"uid"`
			ShortId     string `json:"short_id"`
			Nickname    string `json:"nickname"`
			Signature   string `json:"signature"`
			AvatarThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_medium"`
			FollowStatus    int    `json:"follow_status"`
			IsBlock         bool   `json:"is_block"`
			CustomVerify    string `json:"custom_verify"`
			UniqueId        string `json:"unique_id"`
			RoomId          int    `json:"room_id"`
			AuthorityStatus int    `json:"authority_status"`
			VerifyInfo      string `json:"verify_info"`
			ShareInfo       struct {
				ShareUrl       string `json:"share_url"`
				ShareDesc      string `json:"share_desc"`
				ShareTitle     string `json:"share_title"`
				ShareQrcodeUrl struct {
					Uri       string        `json:"uri"`
					UrlList   []interface{} `json:"url_list"`
					Width     int           `json:"width"`
					Height    int           `json:"height"`
					UrlPrefix interface{}   `json:"url_prefix"`
				} `json:"share_qrcode_url"`
				ShareTitleMyself           string      `json:"share_title_myself"`
				ShareTitleOther            string      `json:"share_title_other"`
				ShareDescInfo              string      `json:"share_desc_info"`
				NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
			} `json:"share_info"`
			WithCommerceEntry      bool        `json:"with_commerce_entry"`
			VerificationType       int         `json:"verification_type"`
			EnterpriseVerifyReason string      `json:"enterprise_verify_reason"`
			IsAdFake               bool        `json:"is_ad_fake"`
			FollowersDetail        interface{} `json:"followers_detail"`
			Region                 string      `json:"region"`
			CommerceUserLevel      int         `json:"commerce_user_level"`
			PlatformSyncInfo       interface{} `json:"platform_sync_info"`
			IsDisciplineMember     bool        `json:"is_discipline_member"`
			Secret                 int         `json:"secret"`
			PreventDownload        bool        `json:"prevent_download"`
			Geofencing             interface{} `json:"geofencing"`
			VideoIcon              struct {
				Uri       string        `json:"uri"`
				UrlList   []interface{} `json:"url_list"`
				Width     int           `json:"width"`
				Height    int           `json:"height"`
				UrlPrefix interface{}   `json:"url_prefix"`
			} `json:"video_icon"`
			FollowerStatus  int `json:"follower_status"`
			CommentSetting  int `json:"comment_setting"`
			DuetSetting     int `json:"duet_setting"`
			DownloadSetting int `json:"download_setting"`
			CoverUrl        []struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_url"`
			Language            string        `json:"language"`
			ItemList            interface{}   `json:"item_list"`
			IsStar              bool          `json:"is_star"`
			TypeLabel           []interface{} `json:"type_label"`
			AdCoverUrl          interface{}   `json:"ad_cover_url"`
			CommentFilterStatus int           `json:"comment_filter_status"`
			Avatar168X168       struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_168x168"`
			Avatar300X300 struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_300x300"`
			RelativeUsers         interface{} `json:"relative_users"`
			ChaList               interface{} `json:"cha_list"`
			SecUid                string      `json:"sec_uid"`
			NeedPoints            interface{} `json:"need_points"`
			HomepageBottomToast   interface{} `json:"homepage_bottom_toast"`
			CanSetGeofencing      interface{} `json:"can_set_geofencing"`
			WhiteCoverUrl         interface{} `json:"white_cover_url"`
			UserTags              interface{} `json:"user_tags"`
			BoldFields            interface{} `json:"bold_fields"`
			SearchHighlight       interface{} `json:"search_highlight"`
			MutualRelationAvatars interface{} `json:"mutual_relation_avatars"`
			Events                interface{} `json:"events"`
			MatchedFriend         struct {
				VideoItems interface{} `json:"video_items"`
			} `json:"matched_friend"`
			AdvanceFeatureItemOrder    interface{} `json:"advance_feature_item_order"`
			AdvancedFeatureInfo        interface{} `json:"advanced_feature_info"`
			UserProfileGuide           interface{} `json:"user_profile_guide"`
			ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
			CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
			AccountLabels              interface{} `json:"account_labels"`
			SocialInfo                 string      `json:"social_info,omitempty"`
		} `json:"author"`
		Music struct {
			Id         int64  `json:"id"`
			IdStr      string `json:"id_str"`
			Title      string `json:"title"`
			Author     string `json:"author"`
			Album      string `json:"album"`
			CoverLarge struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_large"`
			CoverMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_thumb"`
			PlayUrl struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_url"`
			SourcePlatform    int           `json:"source_platform"`
			Duration          int           `json:"duration"`
			Extra             string        `json:"extra"`
			UserCount         int           `json:"user_count"`
			Position          interface{}   `json:"position"`
			CollectStat       int           `json:"collect_stat"`
			Status            int           `json:"status"`
			OfflineDesc       string        `json:"offline_desc"`
			OwnerId           string        `json:"owner_id,omitempty"`
			OwnerNickname     string        `json:"owner_nickname"`
			IsOriginal        bool          `json:"is_original"`
			Mid               string        `json:"mid"`
			BindedChallengeId int           `json:"binded_challenge_id"`
			AuthorDeleted     bool          `json:"author_deleted"`
			OwnerHandle       string        `json:"owner_handle"`
			AuthorPosition    interface{}   `json:"author_position"`
			PreventDownload   bool          `json:"prevent_download"`
			ExternalSongInfo  []interface{} `json:"external_song_info"`
			SecUid            string        `json:"sec_uid,omitempty"`
			AvatarThumb       struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_medium"`
			PreviewStartTime     float64     `json:"preview_start_time"`
			PreviewEndTime       int         `json:"preview_end_time"`
			IsCommerceMusic      bool        `json:"is_commerce_music"`
			IsOriginalSound      bool        `json:"is_original_sound"`
			Artists              interface{} `json:"artists"`
			LyricShortPosition   interface{} `json:"lyric_short_position"`
			MuteShare            bool        `json:"mute_share"`
			TagList              interface{} `json:"tag_list"`
			IsAuthorArtist       bool        `json:"is_author_artist"`
			IsPgc                bool        `json:"is_pgc"`
			SearchHighlight      interface{} `json:"search_highlight"`
			MultiBitRatePlayInfo interface{} `json:"multi_bit_rate_play_info"`
			TtToDspSongInfos     interface{} `json:"tt_to_dsp_song_infos"`
			RecommendStatus      int         `json:"recommend_status"`
			UncertArtists        interface{} `json:"uncert_artists"`
			StrongBeatUrl        struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"strong_beat_url,omitempty"`
		} `json:"music"`
		ChaList []struct {
			Cid     string `json:"cid"`
			ChaName string `json:"cha_name"`
			Desc    string `json:"desc"`
			Schema  string `json:"schema"`
			Author  struct {
				FollowersDetail            interface{} `json:"followers_detail"`
				PlatformSyncInfo           interface{} `json:"platform_sync_info"`
				Geofencing                 interface{} `json:"geofencing"`
				CoverUrl                   interface{} `json:"cover_url"`
				ItemList                   interface{} `json:"item_list"`
				TypeLabel                  interface{} `json:"type_label"`
				AdCoverUrl                 interface{} `json:"ad_cover_url"`
				RelativeUsers              interface{} `json:"relative_users"`
				ChaList                    interface{} `json:"cha_list"`
				NeedPoints                 interface{} `json:"need_points"`
				HomepageBottomToast        interface{} `json:"homepage_bottom_toast"`
				CanSetGeofencing           interface{} `json:"can_set_geofencing"`
				WhiteCoverUrl              interface{} `json:"white_cover_url"`
				UserTags                   interface{} `json:"user_tags"`
				BoldFields                 interface{} `json:"bold_fields"`
				SearchHighlight            interface{} `json:"search_highlight"`
				MutualRelationAvatars      interface{} `json:"mutual_relation_avatars"`
				Events                     interface{} `json:"events"`
				AdvanceFeatureItemOrder    interface{} `json:"advance_feature_item_order"`
				AdvancedFeatureInfo        interface{} `json:"advanced_feature_info"`
				UserProfileGuide           interface{} `json:"user_profile_guide"`
				ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
				CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
				AccountLabels              interface{} `json:"account_labels"`
			} `json:"author"`
			UserCount int `json:"user_count"`
			ShareInfo struct {
				ShareUrl                   string      `json:"share_url"`
				ShareDesc                  string      `json:"share_desc"`
				ShareTitle                 string      `json:"share_title"`
				BoolPersist                int         `json:"bool_persist"`
				ShareTitleMyself           string      `json:"share_title_myself"`
				ShareTitleOther            string      `json:"share_title_other"`
				ShareSignatureUrl          string      `json:"share_signature_url"`
				ShareSignatureDesc         string      `json:"share_signature_desc"`
				ShareQuote                 string      `json:"share_quote"`
				ShareDescInfo              string      `json:"share_desc_info"`
				NowInvitationCardImageUrls interface{} `json:"now_invitation_card_image_urls"`
			} `json:"share_info"`
			ConnectMusic    []interface{} `json:"connect_music"`
			Type            int           `json:"type"`
			SubType         int           `json:"sub_type"`
			IsPgcshow       bool          `json:"is_pgcshow"`
			CollectStat     int           `json:"collect_stat"`
			IsChallenge     int           `json:"is_challenge"`
			ViewCount       int           `json:"view_count"`
			IsCommerce      bool          `json:"is_commerce"`
			HashtagProfile  string        `json:"hashtag_profile"`
			ChaAttrs        interface{}   `json:"cha_attrs"`
			BannerList      interface{}   `json:"banner_list"`
			ShowItems       interface{}   `json:"show_items"`
			SearchHighlight interface{}   `json:"search_highlight"`
		} `json:"cha_list"`
		Video struct {
			PlayAddr struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlKey    string      `json:"url_key"`
				DataSize  int         `json:"data_size"`
				FileHash  string      `json:"file_hash"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_addr"`
			Cover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover"`
			Height       int `json:"height"`
			Width        int `json:"width"`
			DynamicCover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"dynamic_cover"`
			OriginCover struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"origin_cover"`
			Ratio        string `json:"ratio"`
			DownloadAddr struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				DataSize  int         `json:"data_size"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"download_addr"`
			HasWatermark bool `json:"has_watermark"`
			BitRate      []struct {
				GearName    string `json:"gear_name"`
				QualityType int    `json:"quality_type"`
				BitRate     int    `json:"bit_rate"`
				PlayAddr    struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlKey    string      `json:"url_key"`
					DataSize  int         `json:"data_size"`
					FileHash  string      `json:"file_hash"`
					UrlPrefix interface{} `json:"url_prefix"`
				} `json:"play_addr"`
				IsBytevc1 int         `json:"is_bytevc1"`
				DubInfos  interface{} `json:"dub_infos"`
				HDRType   string      `json:"HDR_type"`
				HDRBit    string      `json:"HDR_bit"`
			} `json:"bit_rate"`
			Duration     int `json:"duration"`
			PlayAddrH264 struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlKey    string      `json:"url_key"`
				DataSize  int         `json:"data_size"`
				FileHash  string      `json:"file_hash"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_addr_h264,omitempty"`
			CdnUrlExpired     int         `json:"cdn_url_expired"`
			NeedSetToken      bool        `json:"need_set_token"`
			CoverTsp          float64     `json:"CoverTsp"`
			MiscDownloadAddrs string      `json:"misc_download_addrs,omitempty"`
			Tags              interface{} `json:"tags"`
			BigThumbs         interface{} `json:"big_thumbs"`
			PlayAddrBytevc1   struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlKey    string      `json:"url_key"`
				DataSize  int         `json:"data_size"`
				FileHash  string      `json:"file_hash"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_addr_bytevc1,omitempty"`
			IsBytevc1     int    `json:"is_bytevc1"`
			Meta          string `json:"meta"`
			CoverIsCustom bool   `json:"cover_is_custom,omitempty"`
			ClaInfo       struct {
				HasOriginalAudio     int `json:"has_original_audio"`
				EnableAutoCaption    int `json:"enable_auto_caption"`
				OriginalLanguageInfo struct {
					Lang                                         string `json:"lang"`
					LanguageId                                   int    `json:"language_id"`
					LanguageCode                                 string `json:"language_code"`
					CanTranslateRealtime                         bool   `json:"can_translate_realtime"`
					OriginalCaptionType                          int    `json:"original_caption_type"`
					IsBurninCaption                              bool   `json:"is_burnin_caption"`
					CanTranslateRealtimeSkipTranslationLangCheck bool   `json:"can_translate_realtime_skip_translation_lang_check"`
				} `json:"original_language_info"`
				CaptionInfos []struct {
					Lang              string   `json:"lang"`
					LanguageId        int      `json:"language_id"`
					Url               string   `json:"url"`
					Expire            int      `json:"expire"`
					CaptionFormat     string   `json:"caption_format"`
					ComplaintId       int64    `json:"complaint_id"`
					IsAutoGenerated   bool     `json:"is_auto_generated"`
					SubId             int      `json:"sub_id"`
					SubVersion        string   `json:"sub_version"`
					ClaSubtitleId     int64    `json:"cla_subtitle_id"`
					TranslatorId      int      `json:"translator_id"`
					LanguageCode      string   `json:"language_code"`
					IsOriginalCaption bool     `json:"is_original_caption"`
					UrlList           []string `json:"url_list"`
					CaptionLength     int      `json:"caption_length"`
				} `json:"caption_infos"`
				CreatorEditedCaptionId int         `json:"creator_edited_caption_id"`
				VerticalPositions      interface{} `json:"vertical_positions"`
				HideOriginalCaption    bool        `json:"hide_original_caption"`
				CaptionsType           int         `json:"captions_type"`
				NoCaptionReason        int         `json:"no_caption_reason"`
			} `json:"cla_info,omitempty"`
			SourceHDRType int           `json:"source_HDR_type"`
			BitRateAudio  []interface{} `json:"bit_rate_audio"`
		} `json:"video"`
		ShareUrl   string `json:"share_url"`
		UserDigged int    `json:"user_digged"`
		Statistics struct {
			AwemeId            string `json:"aweme_id"`
			CommentCount       int    `json:"comment_count"`
			DiggCount          int    `json:"digg_count"`
			DownloadCount      int    `json:"download_count"`
			PlayCount          int    `json:"play_count"`
			ShareCount         int    `json:"share_count"`
			ForwardCount       int    `json:"forward_count"`
			LoseCount          int    `json:"lose_count"`
			LoseCommentCount   int    `json:"lose_comment_count"`
			WhatsappShareCount int    `json:"whatsapp_share_count"`
			CollectCount       int    `json:"collect_count"`
		} `json:"statistics"`
		Status struct {
			AwemeId        string `json:"aweme_id"`
			IsDelete       bool   `json:"is_delete"`
			AllowShare     bool   `json:"allow_share"`
			AllowComment   bool   `json:"allow_comment"`
			IsPrivate      bool   `json:"is_private"`
			WithGoods      bool   `json:"with_goods"`
			PrivateStatus  int    `json:"private_status"`
			InReviewing    bool   `json:"in_reviewing"`
			Reviewed       int    `json:"reviewed"`
			SelfSee        bool   `json:"self_see"`
			IsProhibited   bool   `json:"is_prohibited"`
			DownloadStatus int    `json:"download_status"`
		} `json:"status"`
		Rate      int `json:"rate"`
		TextExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			UserId      string `json:"user_id"`
			Type        int    `json:"type"`
			SecUid      string `json:"sec_uid"`
			SubType     int    `json:"sub_type,omitempty"`
			HashtagName string `json:"hashtag_name,omitempty"`
			HashtagId   string `json:"hashtag_id,omitempty"`
			IsCommerce  bool   `json:"is_commerce,omitempty"`
		} `json:"text_extra"`
		IsTop    int `json:"is_top"`
		LabelTop struct {
			Uri       string      `json:"uri"`
			UrlList   []string    `json:"url_list"`
			Width     int         `json:"width"`
			Height    int         `json:"height"`
			UrlPrefix interface{} `json:"url_prefix"`
		} `json:"label_top"`
		ShareInfo struct {
			ShareUrl                    string      `json:"share_url"`
			ShareDesc                   string      `json:"share_desc"`
			ShareTitle                  string      `json:"share_title"`
			BoolPersist                 int         `json:"bool_persist"`
			ShareTitleMyself            string      `json:"share_title_myself"`
			ShareTitleOther             string      `json:"share_title_other"`
			ShareLinkDesc               string      `json:"share_link_desc"`
			ShareSignatureUrl           string      `json:"share_signature_url"`
			ShareSignatureDesc          string      `json:"share_signature_desc"`
			ShareQuote                  string      `json:"share_quote"`
			WhatsappDesc                string      `json:"whatsapp_desc"`
			ShareDescInfo               string      `json:"share_desc_info"`
			NowInvitationCardImageUrls  interface{} `json:"now_invitation_card_image_urls"`
			ShareButtonDisplayMode      int         `json:"share_button_display_mode"`
			ButtonDisplayStrategeSource string      `json:"button_display_stratege_source,omitempty"`
		} `json:"share_info"`
		Distance    string        `json:"distance"`
		VideoLabels []interface{} `json:"video_labels"`
		IsVr        bool          `json:"is_vr"`
		IsAds       bool          `json:"is_ads"`
		AwemeType   int           `json:"aweme_type"`
		CmtSwt      bool          `json:"cmt_swt"`
		ImageInfos  interface{}   `json:"image_infos"`
		RiskInfos   struct {
			Vote     bool   `json:"vote"`
			Warn     bool   `json:"warn"`
			RiskSink bool   `json:"risk_sink"`
			Type     int    `json:"type"`
			Content  string `json:"content"`
		} `json:"risk_infos"`
		IsRelieve            bool          `json:"is_relieve"`
		SortLabel            string        `json:"sort_label"`
		Position             interface{}   `json:"position"`
		UniqidPosition       interface{}   `json:"uniqid_position"`
		AuthorUserId         int64         `json:"author_user_id"`
		BodydanceScore       int           `json:"bodydance_score"`
		Geofencing           interface{}   `json:"geofencing"`
		IsHashTag            int           `json:"is_hash_tag"`
		IsPgcshow            bool          `json:"is_pgcshow"`
		Region               string        `json:"region"`
		VideoText            []interface{} `json:"video_text"`
		CollectStat          int           `json:"collect_stat"`
		LabelTopText         interface{}   `json:"label_top_text"`
		GroupId              string        `json:"group_id"`
		PreventDownload      bool          `json:"prevent_download"`
		NicknamePosition     interface{}   `json:"nickname_position"`
		ChallengePosition    interface{}   `json:"challenge_position"`
		ItemCommentSettings  int           `json:"item_comment_settings"`
		WithPromotionalMusic bool          `json:"with_promotional_music"`
		LongVideo            interface{}   `json:"long_video"`
		ItemDuet             int           `json:"item_duet"`
		ItemReact            int           `json:"item_react"`
		DescLanguage         string        `json:"desc_language"`
		InteractionStickers  []struct {
			Type      int    `json:"type"`
			Index     int    `json:"index"`
			TrackInfo string `json:"track_info"`
			Attr      string `json:"attr"`
		} `json:"interaction_stickers"`
		MiscInfo           string      `json:"misc_info"`
		OriginCommentIds   interface{} `json:"origin_comment_ids"`
		CommerceConfigData interface{} `json:"commerce_config_data"`
		DistributeType     int         `json:"distribute_type"`
		VideoControl       struct {
			AllowDownload         bool `json:"allow_download"`
			ShareType             int  `json:"share_type"`
			ShowProgressBar       int  `json:"show_progress_bar"`
			DraftProgressBar      int  `json:"draft_progress_bar"`
			AllowDuet             bool `json:"allow_duet"`
			AllowReact            bool `json:"allow_react"`
			PreventDownloadType   int  `json:"prevent_download_type"`
			AllowDynamicWallpaper bool `json:"allow_dynamic_wallpaper"`
			TimerStatus           int  `json:"timer_status"`
			AllowMusic            bool `json:"allow_music"`
			AllowStitch           bool `json:"allow_stitch"`
		} `json:"video_control"`
		HasVsEntry   bool `json:"has_vs_entry"`
		CommerceInfo struct {
			AdvPromotable      bool `json:"adv_promotable"`
			BrandedContentType int  `json:"branded_content_type"`
		} `json:"commerce_info"`
		NeedVsEntry    bool `json:"need_vs_entry"`
		VideoReplyInfo struct {
			AwemeId        int64 `json:"aweme_id"`
			CommentId      int64 `json:"comment_id"`
			AliasCommentId int64 `json:"alias_comment_id"`
		} `json:"video_reply_info,omitempty"`
		Anchors           interface{} `json:"anchors"`
		HybridLabel       interface{} `json:"hybrid_label"`
		WithSurvey        bool        `json:"with_survey"`
		GeofencingRegions interface{} `json:"geofencing_regions"`
		AwemeAcl          struct {
			DownloadGeneral struct {
				Code      int    `json:"code"`
				ShowType  int    `json:"show_type"`
				Transcode int    `json:"transcode"`
				Mute      bool   `json:"mute"`
				Extra     string `json:"extra,omitempty"`
			} `json:"download_general"`
			DownloadMaskPanel struct {
				Code      int    `json:"code"`
				ShowType  int    `json:"show_type"`
				Transcode int    `json:"transcode"`
				Mute      bool   `json:"mute"`
				Extra     string `json:"extra,omitempty"`
			} `json:"download_mask_panel"`
			ShareListStatus int `json:"share_list_status"`
			ShareGeneral    struct {
				Code      int    `json:"code"`
				ShowType  int    `json:"show_type"`
				Transcode int    `json:"transcode"`
				Mute      bool   `json:"mute"`
				ToastMsg  string `json:"toast_msg,omitempty"`
				Extra     string `json:"extra,omitempty"`
			} `json:"share_general"`
			PlatformList    interface{} `json:"platform_list"`
			ShareActionList interface{} `json:"share_action_list"`
			PressActionList interface{} `json:"press_action_list"`
		} `json:"aweme_acl"`
		CoverLabels          interface{}   `json:"cover_labels"`
		MaskInfos            []interface{} `json:"mask_infos"`
		SearchHighlight      interface{}   `json:"search_highlight"`
		PlaylistBlocked      bool          `json:"playlist_blocked"`
		GreenScreenMaterials interface{}   `json:"green_screen_materials"`
		InteractPermission   struct {
			Duet                 int `json:"duet"`
			Stitch               int `json:"stitch"`
			DuetPrivacySetting   int `json:"duet_privacy_setting"`
			StitchPrivacySetting int `json:"stitch_privacy_setting"`
			Upvote               int `json:"upvote"`
			AllowAddingToStory   int `json:"allow_adding_to_story"`
			AllowCreateSticker   struct {
				Status int `json:"status"`
			} `json:"allow_create_sticker"`
		} `json:"interact_permission"`
		QuestionList     interface{} `json:"question_list"`
		ContentDesc      string      `json:"content_desc"`
		ContentDescExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
			HashtagId   string `json:"hashtag_id"`
			IsCommerce  bool   `json:"is_commerce"`
			LineIdx     int    `json:"line_idx"`
		} `json:"content_desc_extra"`
		ProductsInfo              interface{} `json:"products_info"`
		FollowUpPublishFromId     int         `json:"follow_up_publish_from_id"`
		DisableSearchTrendingBar  bool        `json:"disable_search_trending_bar"`
		MusicBeginTimeInMs        int         `json:"music_begin_time_in_ms"`
		MusicEndTimeInMs          int         `json:"music_end_time_in_ms,omitempty"`
		ItemDistributeSource      string      `json:"item_distribute_source"`
		ItemSourceCategory        int         `json:"item_source_category"`
		BrandedContentAccounts    interface{} `json:"branded_content_accounts"`
		IsDescriptionTranslatable bool        `json:"is_description_translatable"`
		FollowUpItemIdGroups      string      `json:"follow_up_item_id_groups"`
		IsTextStickerTranslatable bool        `json:"is_text_sticker_translatable"`
		TextStickerMajorLang      string      `json:"text_sticker_major_lang"`
		OriginalClientText        struct {
			MarkupText string `json:"markup_text"`
			TextExtra  []struct {
				Start       int    `json:"start,omitempty"`
				End         int    `json:"end,omitempty"`
				UserId      string `json:"user_id,omitempty"`
				Type        int    `json:"type"`
				IsCommerce  bool   `json:"is_commerce,omitempty"`
				SubType     int    `json:"sub_type,omitempty"`
				LineIdx     int    `json:"line_idx,omitempty"`
				TagId       string `json:"tag_id"`
				HashtagName string `json:"hashtag_name,omitempty"`
			} `json:"text_extra"`
		} `json:"original_client_text"`
		MusicSelectedFrom       string      `json:"music_selected_from"`
		TtsVoiceIds             interface{} `json:"tts_voice_ids"`
		ReferenceTtsVoiceIds    interface{} `json:"reference_tts_voice_ids"`
		VoiceFilterIds          interface{} `json:"voice_filter_ids"`
		ReferenceVoiceFilterIds interface{} `json:"reference_voice_filter_ids"`
		MusicTitleStyle         int         `json:"music_title_style"`
		CommentConfig           struct {
			EmojiRecommendList interface{} `json:"emoji_recommend_list"`
		} `json:"comment_config"`
		AddedSoundMusicInfo struct {
			Id         int64  `json:"id"`
			IdStr      string `json:"id_str"`
			Title      string `json:"title"`
			Author     string `json:"author"`
			Album      string `json:"album"`
			CoverLarge struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_large"`
			CoverMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"cover_thumb"`
			PlayUrl struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"play_url"`
			SourcePlatform    int           `json:"source_platform"`
			Duration          int           `json:"duration"`
			Extra             string        `json:"extra"`
			UserCount         int           `json:"user_count"`
			Position          interface{}   `json:"position"`
			CollectStat       int           `json:"collect_stat"`
			Status            int           `json:"status"`
			OfflineDesc       string        `json:"offline_desc"`
			OwnerId           string        `json:"owner_id,omitempty"`
			OwnerNickname     string        `json:"owner_nickname"`
			IsOriginal        bool          `json:"is_original"`
			Mid               string        `json:"mid"`
			BindedChallengeId int           `json:"binded_challenge_id"`
			AuthorDeleted     bool          `json:"author_deleted"`
			OwnerHandle       string        `json:"owner_handle"`
			AuthorPosition    interface{}   `json:"author_position"`
			PreventDownload   bool          `json:"prevent_download"`
			ExternalSongInfo  []interface{} `json:"external_song_info"`
			SecUid            string        `json:"sec_uid,omitempty"`
			AvatarThumb       struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_medium"`
			PreviewStartTime     float64     `json:"preview_start_time"`
			PreviewEndTime       int         `json:"preview_end_time"`
			IsCommerceMusic      bool        `json:"is_commerce_music"`
			IsOriginalSound      bool        `json:"is_original_sound"`
			Artists              interface{} `json:"artists"`
			LyricShortPosition   interface{} `json:"lyric_short_position"`
			MuteShare            bool        `json:"mute_share"`
			TagList              interface{} `json:"tag_list"`
			IsAuthorArtist       bool        `json:"is_author_artist"`
			IsPgc                bool        `json:"is_pgc"`
			SearchHighlight      interface{} `json:"search_highlight"`
			MultiBitRatePlayInfo interface{} `json:"multi_bit_rate_play_info"`
			TtToDspSongInfos     interface{} `json:"tt_to_dsp_song_infos"`
			RecommendStatus      int         `json:"recommend_status"`
			UncertArtists        interface{} `json:"uncert_artists"`
			StrongBeatUrl        struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"strong_beat_url,omitempty"`
		} `json:"added_sound_music_info"`
		OriginVolume               string      `json:"origin_volume"`
		MusicVolume                string      `json:"music_volume"`
		SupportDanmaku             bool        `json:"support_danmaku"`
		HasDanmaku                 bool        `json:"has_danmaku"`
		MufCommentInfoV2           interface{} `json:"muf_comment_info_v2"`
		BehindTheSongMusicIds      interface{} `json:"behind_the_song_music_ids"`
		BehindTheSongVideoMusicIds interface{} `json:"behind_the_song_video_music_ids"`
		ContentOriginalType        int         `json:"content_original_type"`
		OperatorBoostInfo          interface{} `json:"operator_boost_info"`
		LogInfo                    struct {
			Order string `json:"order"`
		} `json:"log_info"`
		MainArchCommon string `json:"main_arch_common"`
		AigcInfo       struct {
			AigcLabelType int `json:"aigc_label_type"`
		} `json:"aigc_info"`
		Banners           interface{} `json:"banners"`
		PickedUsers       interface{} `json:"picked_users"`
		CommentTopbarInfo interface{} `json:"comment_topbar_info"`
		PlaylistInfo      struct {
			MixId     string `json:"mix_id"`
			Name      string `json:"name"`
			Index     int    `json:"index"`
			ItemTotal int    `json:"item_total"`
		} `json:"playlist_info,omitempty"`
		ShootTabName       string `json:"shoot_tab_name,omitempty"`
		ContentType        string `json:"content_type,omitempty"`
		ContentSizeType    int    `json:"content_size_type,omitempty"`
		BatchIndex         int    `json:"batch_index,omitempty"`
		InteractionTagInfo struct {
			InterestLevel  int    `json:"interest_level"`
			VideoLabelText string `json:"video_label_text"`
			TaggedUsers    []struct {
				Uid           string `json:"uid"`
				UniqueId      string `json:"unique_id"`
				Nickname      string `json:"nickname"`
				Avatar168X168 struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlPrefix interface{} `json:"url_prefix"`
				} `json:"avatar_168x168"`
				AvatarThumb struct {
					Uri       string      `json:"uri"`
					UrlList   []string    `json:"url_list"`
					Width     int         `json:"width"`
					Height    int         `json:"height"`
					UrlPrefix interface{} `json:"url_prefix"`
				} `json:"avatar_thumb"`
				FollowStatus      int    `json:"follow_status"`
				FollowerStatus    int    `json:"follower_status"`
				CustomVerify      string `json:"custom_verify"`
				InterestLevel     int    `json:"interest_level"`
				IsBusinessAccount bool   `json:"is_business_account"`
				InvitationStatus  int    `json:"invitation_status"`
			} `json:"tagged_users"`
		} `json:"interaction_tag_info,omitempty"`
	} `json:"aweme_list"`
	HomeModel    int `json:"home_model"`
	RefreshClear int `json:"refresh_clear"`
	Extra        struct {
		Now          int64       `json:"now"`
		FatalItemIds interface{} `json:"fatal_item_ids"`
		ApiDebugInfo interface{} `json:"api_debug_info"`
	} `json:"extra"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	PreloadAds    []interface{} `json:"preload_ads"`
	PreloadAwemes interface{}   `json:"preload_awemes"`
	LogInfo       struct {
		ImprId   string `json:"impr_id"`
		PullType string `json:"pull_type"`
	} `json:"log_info"`
}

// New returns a tiktok extractor.
func New() extractors.Extractor {
	return &extractor{}
}

func getTiktokVideoId(url string) (string, string) {

	validURL := `https?://www\.tiktok\.com/(?:embed|@(?P<user_id>[\w\.-]+)?/video)/(?P<id>\d+)`
	re := regexp.MustCompile(validURL)
	match := re.FindStringSubmatch(url)

	var videoId, userId string

	for i, n := range re.SubexpNames() {
		if i != 0 && n != "" {
			//if n == "user_id" {
			//	userId = match[i]
			//}
			if n == "id" {
				videoId = match[i]
			}
		}
	}

	return userId, videoId
}

func getRealUrl(url string) (string, error) {

	client := http.Client{
		Transport: &http.Transport{
			Proxy:              http.ProxyFromEnvironment,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求出错: %v", err)
	}

	req.Header.Set("User-Agent", "facebookexternalhit/1.1")

	//if _, ok := headers["Referer"]; !ok {
	//	req.Header.Set("Referer", urlStr)
	//}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求出错: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusMovedPermanently &&
		resp.StatusCode != http.StatusFound && resp.StatusCode != http.StatusSeeOther &&
		resp.StatusCode != http.StatusTemporaryRedirect {

		return "", fmt.Errorf("请求返回非重定向状态码: %v", resp.StatusCode)
	}

	url = resp.Request.URL.String()

	return url, nil
}

func getRealUrlBrowser(url string) (string, error) {

	client := http.Client{
		Transport: &http.Transport{
			Proxy:              http.ProxyFromEnvironment,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求出错: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	//if _, ok := headers["Referer"]; !ok {
	//	req.Header.Set("Referer", urlStr)
	//}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求出错: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusMovedPermanently &&
		resp.StatusCode != http.StatusFound && resp.StatusCode != http.StatusSeeOther &&
		resp.StatusCode != http.StatusTemporaryRedirect {

		return "", fmt.Errorf("请求返回非重定向状态码: %v", resp.StatusCode)
	}

	url = resp.Request.URL.String()

	return url, nil
}

func randOpenid() string {
	rand.Seed(time.Now().UnixNano())

	chars := "0123456789abcdef"
	length := 16

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	randomString := string(result)

	return randomString
}

func randUuid() string {
	rand.Seed(time.Now().UnixNano())

	chars := "0123456789"
	length := 16

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	randomString := string(result)

	return randomString
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {

	if !strings.Contains(url, "/@") {
		url, _ = getRealUrl(url)
	}

	_, videoId := getTiktokVideoId(url)

	if len(videoId) <= 0 {
		return nil, errors.New("无法解析视频ID")
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 创建 GET 请求
	req, err := http.NewRequest("GET", "https://api22-normal-c-useast2a.tiktokv.com/aweme/v1/feed/?aweme_id="+videoId+"&version_name=26.1.3&version_code=260103&build_number=26.1.3&manifest_version_code=260103&update_version_code=260103&openudid="+randOpenid()+"&uuid="+randUuid()+"&_rticket="+strconv.FormatInt(time.Now().Unix()*1000, 10)+"&ts="+strconv.FormatInt(time.Now().Unix(), 10)+"&device_brand=Google&device_type=Pixel+7&device_platform=android&resolution=1080%2A2400&dpi=420&os_version=13&os_api=29&carrier_region=US&sys_region=US&region=US&app_name=trill&app_language=en&language=en&timezone_name=America%2FNew_York&timezone_offset=-14400&channel=googleplay&ac=wifi&mcc_mnc=310260&is_my_cn=0&aid=1180&ssmix=a&as=a1qwert123&cp=cbfhckdckkde1", nil)
	if err != nil {
		return nil, errors2.Annotate(err, "创建请求失败")
	}

	// 设置请求的 User-Agent 和 Accept 头部
	req.Header.Set("User-Agent", "com.ss.android.ugc.trill/260103 (Linux; U; Android 13; en_US; Pixel 7; Build/TD1A.220804.031; Cronet/58.0.2991.0)")
	req.Header.Set("Accept", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors2.Annotate(err, "请求发送失败")
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	var result tiktokContent
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors2.Annotate(err, "json 解析失败")
	}

	var extractorsData extractors.Data

	for _, item := range result.AwemeList {
		if item.AwemeId == videoId {
			extractorsData.Title = item.Desc
			extractorsData.Cover = item.Video.Cover.UrlList[0]
			extractorsData.IsNotVideo = false
			extractorsData.Type = extractors.DataTypeVideo
			extractorsData.Site = "TikTok tiktok.com"
			extractorsData.URL = url
			for _, u := range item.Video.PlayAddr.UrlList {

				if strings.Contains(u, "tiktokv.com") {
					url, err = getRealUrl(u)
					if err == nil {
						extractorsData.Streams = map[string]*extractors.Stream{
							"default": {
								Parts: []*extractors.Part{
									{
										URL:  url,
										Ext:  "mp4",
										Size: 850,
									},
								},
								Size:    850,
								Quality: "1080p",
							},
						}
						break
					}
				}
			}

		}
	}

	return []*extractors.Data{&extractorsData}, nil
}
