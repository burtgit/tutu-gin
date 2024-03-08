package facebook

import (
	"encoding/json"
	"fmt"
	errors2 "github.com/juju/errors"
	"net/http"
	"strings"
	"tutu-gin/lib/extractors"

	"github.com/pkg/errors"

	"github.com/iawia002/lux/request"
	"github.com/iawia002/lux/utils"
)

func init() {
	extractors.Register("facebook", New())
}

type extractor struct{}

// New returns a facebook extractor.
func New() extractors.Extractor {
	return &extractor{}
}

type facebook struct {
	Bbox struct {
		Complete bool `json:"complete"`
		Result   struct {
			Data struct {
				Video struct {
					Story struct {
						Attachments []struct {
							Media struct {
								Typename           string `json:"__typename"`
								PreferredThumbnail struct {
									Image struct {
										Uri string `json:"uri"`
									} `json:"image"`
									ImagePreviewPayload interface{} `json:"image_preview_payload"`
									Id                  string      `json:"id"`
								} `json:"preferred_thumbnail"`
								Id                                     string        `json:"id"`
								AnimatedImageCaption                   interface{}   `json:"animated_image_caption"`
								Width                                  int           `json:"width"`
								Height                                 int           `json:"height"`
								BroadcasterOrigin                      interface{}   `json:"broadcaster_origin"`
								BroadcastId                            interface{}   `json:"broadcast_id"`
								BroadcastStatus                        interface{}   `json:"broadcast_status"`
								IsLiveStreaming                        bool          `json:"is_live_streaming"`
								IsLiveTraceEnabled                     bool          `json:"is_live_trace_enabled"`
								IsLooping                              bool          `json:"is_looping"`
								IsVideoBroadcast                       bool          `json:"is_video_broadcast"`
								IsPodcastVideo                         bool          `json:"is_podcast_video"`
								LoopCount                              int           `json:"loop_count"`
								IsSpherical                            bool          `json:"is_spherical"`
								IsSphericalEnabled                     bool          `json:"is_spherical_enabled"`
								UnsupportedBrowserMessage              interface{}   `json:"unsupported_browser_message"`
								PmvMetadata                            interface{}   `json:"pmv_metadata"`
								LatencySensitiveConfig                 interface{}   `json:"latency_sensitive_config"`
								LivePlaybackInstrumentationConfigs     interface{}   `json:"live_playback_instrumentation_configs"`
								IsNcsr                                 bool          `json:"is_ncsr"`
								PermalinkUrl                           string        `json:"permalink_url"`
								CaptionsUrl                            interface{}   `json:"captions_url"`
								DashPrefetchExperimental               []string      `json:"dash_prefetch_experimental"`
								VideoAvailableCaptionsLocales          []interface{} `json:"video_available_captions_locales"`
								VideoStatusType                        string        `json:"video_status_type"`
								CanUseOz                               bool          `json:"can_use_oz"`
								DashManifest                           string        `json:"dash_manifest"`
								DashManifestUrl                        string        `json:"dash_manifest_url"`
								MinQualityPreference                   string        `json:"min_quality_preference"`
								AudioUserPreferredLanguage             string        `json:"audio_user_preferred_language"`
								IsRssPodcastVideo                      bool          `json:"is_rss_podcast_video"`
								BrowserNativeSdUrl                     string        `json:"browser_native_sd_url"`
								BrowserNativeHdUrl                     string        `json:"browser_native_hd_url"`
								SphericalVideoFallbackUrls             interface{}   `json:"spherical_video_fallback_urls"`
								IsGamingVideo                          bool          `json:"is_gaming_video"`
								IsLatencyMenuEnabled                   bool          `json:"is_latency_menu_enabled"`
								FblsTier                               interface{}   `json:"fbls_tier"`
								IsLatencySensitiveBroadcast            bool          `json:"is_latency_sensitive_broadcast"`
								CometVideoPlayerStaticConfig           string        `json:"comet_video_player_static_config"`
								CometVideoPlayerContextSensitiveConfig string        `json:"comet_video_player_context_sensitive_config"`
								VideoPlayerShakaPerformanceLoggerInit  struct {
									Typename                                                          string `json:"__typename"`
									ModuleOperationUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
									ModuleComponentUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
								} `json:"video_player_shaka_performance_logger_init"`
								VideoPlayerShakaPerformanceLoggerShouldSample bool `json:"video_player_shaka_performance_logger_should_sample"`
								VideoPlayerShakaPerformanceLoggerInit2        struct {
									Typename                                                        string `json:"__typename"`
									ModuleOperationUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
									ModuleComponentUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
									PerSessionSamplingRate interface{} `json:"per_session_sampling_rate"`
								} `json:"video_player_shaka_performance_logger_init2"`
								AutoplayGatingResult      string        `json:"autoplay_gating_result"`
								ViewerAutoplaySetting     string        `json:"viewer_autoplay_setting"`
								CanAutoplay               bool          `json:"can_autoplay"`
								DrmInfo                   string        `json:"drm_info"`
								P2PSettings               interface{}   `json:"p2p_settings"`
								AudioSettings             interface{}   `json:"audio_settings"`
								CaptionsSettings          interface{}   `json:"captions_settings"`
								BroadcastLowLatencyConfig interface{}   `json:"broadcast_low_latency_config"`
								AudioAvailability         string        `json:"audio_availability"`
								MutedSegments             []interface{} `json:"muted_segments"`
								SphericalVideoRenderer    interface{}   `json:"spherical_video_renderer"`
								VideoImfData              interface{}   `json:"video_imf_data"`
								IsClippingEnabled         bool          `json:"is_clipping_enabled"`
								LiveRewindEnabled         bool          `json:"live_rewind_enabled"`
								Owner                     struct {
									Typename                        string `json:"__typename"`
									Id                              string `json:"id"`
									IsVideoOwner                    string `json:"__isVideoOwner"`
									HasProfessionalFeaturesForWatch bool   `json:"has_professional_features_for_watch"`
								} `json:"owner"`
								PlayableDurationInMs                   int         `json:"playable_duration_in_ms"`
								IsHuddle                               bool        `json:"is_huddle"`
								Url                                    string      `json:"url"`
								IfViewerCanUseLatencyMenu              interface{} `json:"if_viewer_can_use_latency_menu"`
								IfViewerCanUseLatencyMenuToggle        interface{} `json:"if_viewer_can_use_latency_menu_toggle"`
								IfViewerCanSeeCommunityModerationTools interface{} `json:"if_viewer_can_see_community_moderation_tools"`
								IfViewerCanUseLiveRewind               interface{} `json:"if_viewer_can_use_live_rewind"`
								IfViewerCanUseClipping                 interface{} `json:"if_viewer_can_use_clipping"`
								IfViewerCanSeeCostreamingTools         interface{} `json:"if_viewer_can_see_costreaming_tools"`
								VideoPlayerScrubberPreviewRenderer     struct {
									Typename string `json:"__typename"`
									Video    struct {
										ScrubberPreviewThumbnailInformation struct {
											SpriteUris                 []string `json:"sprite_uris"`
											ThumbnailWidth             int      `json:"thumbnail_width"`
											ThumbnailHeight            int      `json:"thumbnail_height"`
											HasPreviewThumbnails       bool     `json:"has_preview_thumbnails"`
											NumImagesPerRow            int      `json:"num_images_per_row"`
											MaxNumberOfImagesPerSprite int      `json:"max_number_of_images_per_sprite"`
											TimeIntervalBetweenImage   int      `json:"time_interval_between_image"`
										} `json:"scrubber_preview_thumbnail_information"`
										Id string `json:"id"`
									} `json:"video"`
									ModuleOperationVideoPlayerScrubberPreviewVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_operation_VideoPlayerScrubberPreview_video"`
									ModuleComponentVideoPlayerScrubberPreviewVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_component_VideoPlayerScrubberPreview_video"`
								} `json:"video_player_scrubber_preview_renderer"`
								VideoPlayerScrubberBaseContentRenderer interface{} `json:"video_player_scrubber_base_content_renderer"`
								RecipientGroup                         interface{} `json:"recipient_group"`
								MusicAttachmentMetadata                interface{} `json:"music_attachment_metadata"`
								VideoContainerType                     string      `json:"video_container_type"`
								BreakingStatus                         bool        `json:"breakingStatus"`
								VideoId                                string      `json:"videoId"`
								IsPremiere                             bool        `json:"isPremiere"`
								LiveViewerCount                        int         `json:"liveViewerCount"`
								RehearsalInfo                          interface{} `json:"rehearsalInfo"`
								IsLiveAudioRoomV2Broadcast             bool        `json:"is_live_audio_room_v2_broadcast"`
								PublishTime                            int         `json:"publish_time"`
								LiveSpeakerCountIndicator              interface{} `json:"live_speaker_count_indicator"`
								CometProductTagFeedOverlayRenderer     interface{} `json:"comet_product_tag_feed_overlay_renderer"`
								IfViewerCanSeePayToAccessPaywall       interface{} `json:"if_viewer_can_see_pay_to_access_paywall"`
								IsNode                                 string      `json:"__isNode"`
							} `json:"media"`
						} `json:"attachments"`
						Id string `json:"id"`
					} `json:"story"`
					Id string `json:"id"`
				} `json:"video"`
				VideoHomeWwwMatchaEntities []interface{} `json:"video_home_www_matcha_entities"`
			} `json:"data"`
			Extensions struct {
				AllVideoDashPrefetchRepresentations []struct {
					Representations []struct {
						RepresentationId      string `json:"representation_id"`
						MimeType              string `json:"mime_type"`
						Codecs                string `json:"codecs"`
						BaseUrl               string `json:"base_url"`
						Bandwidth             int    `json:"bandwidth"`
						Height                int    `json:"height"`
						Width                 int    `json:"width"`
						PlaybackResolutionMos string `json:"playback_resolution_mos"`
						Segments              []struct {
							Start int `json:"start"`
							End   int `json:"end"`
						} `json:"segments"`
					} `json:"representations"`
					VideoId int64 `json:"video_id"`
				} `json:"all_video_dash_prefetch_representations"`
				IsFinal bool `json:"is_final"`
			} `json:"extensions"`
		} `json:"result"`
		SequenceNumber int `json:"sequence_number"`
	} `json:"__bbox"`
}

type facebookPoster struct {
	Bbox struct {
		Complete bool `json:"complete"`
		Result   struct {
			Data struct {
				Node struct {
					Typename      string      `json:"__typename"`
					IsFeedUnit    string      `json:"__isFeedUnit"`
					DebugInfo     interface{} `json:"debug_info"`
					Id            string      `json:"id"`
					SponsoredData interface{} `json:"sponsored_data"`
					Feedback      struct {
						AssociatedGroup interface{} `json:"associated_group"`
						Id              string      `json:"id"`
					} `json:"feedback"`
					IsStoryCivic     interface{}   `json:"is_story_civic"`
					MatchedTerms     []interface{} `json:"matched_terms"`
					PostId           string        `json:"post_id"`
					CixScreen        interface{}   `json:"cix_screen"`
					FutureOfFeedInfo struct {
						ShouldReverseMessageAndAttachmentPosition bool   `json:"should_reverse_message_and_attachment_position"`
						ShouldOverlayHeader                       bool   `json:"should_overlay_header"`
						AspectRatioUpdate                         int    `json:"aspect_ratio_update"`
						WebReshareVariant                         string `json:"web_reshare_variant"`
					} `json:"future_of_feed_info"`
					AttachedStory interface{} `json:"attached_story"`
					Bumpers       interface{} `json:"bumpers"`
					CometSections struct {
						Content struct {
							Typename             string `json:"__typename"`
							IsICometStorySection string `json:"__isICometStorySection"`
							IsProdEligible       bool   `json:"is_prod_eligible"`
							Story                struct {
								Feedback struct {
									Id string `json:"id"`
								} `json:"feedback"`
								Id            string `json:"id"`
								CometSections struct {
									ContextLayout struct {
										Typename                      string      `json:"__typename"`
										IsICometStorySection          string      `json:"__isICometStorySection"`
										IsProdEligible                bool        `json:"is_prod_eligible"`
										LocalAlertsStoryMenuPromotion interface{} `json:"local_alerts_story_menu_promotion"`
										Story                         struct {
											Id                        string      `json:"id"`
											DebugInfo                 interface{} `json:"debug_info"`
											SerializedFrtpIdentifiers interface{} `json:"serialized_frtp_identifiers"`
											CanViewerSeeMenu          bool        `json:"can_viewer_see_menu"`
											CometSections             struct {
												ActorPhoto struct {
													Typename             string `json:"__typename"`
													IsICometStorySection string `json:"__isICometStorySection"`
													IsProdEligible       bool   `json:"is_prod_eligible"`
													Story                struct {
														Actors []struct {
															Typename              string      `json:"__typename"`
															IsActor               string      `json:"__isActor"`
															Id                    string      `json:"id"`
															IsEntity              string      `json:"__isEntity"`
															Url                   string      `json:"url"`
															WorkForeignEntityInfo interface{} `json:"work_foreign_entity_info"`
															WorkInfo              interface{} `json:"work_info"`
															StoryBucket           struct {
																Nodes []struct {
																	ShouldShowCloseFriendBadge bool   `json:"should_show_close_friend_badge"`
																	Id                         string `json:"id"`
																	FirstStoryToShow           struct {
																		Id                 string `json:"id"`
																		StoryCardSeenState struct {
																			IsSeenByViewer bool `json:"is_seen_by_viewer"`
																		} `json:"story_card_seen_state"`
																	} `json:"first_story_to_show"`
																} `json:"nodes"`
															} `json:"story_bucket"`
															LiveVideoForCometLiveRing interface{} `json:"live_video_for_comet_live_ring"`
															AnswerAgentGroupId        interface{} `json:"answer_agent_group_id"`
															ProfileUrl                string      `json:"profile_url"`
															Name                      string      `json:"name"`
															ProfilePicture            struct {
																Uri    string `json:"uri"`
																Width  int    `json:"width"`
																Height int    `json:"height"`
																Scale  int    `json:"scale"`
															} `json:"profile_picture"`
															IsAdditionalProfilePlus bool        `json:"is_additional_profile_plus"`
															DelegatePage            interface{} `json:"delegate_page"`
														} `json:"actors"`
														CometSections struct {
															ActionLink interface{} `json:"action_link"`
														} `json:"comet_sections"`
														Attachments []struct {
															ActionLinks []interface{} `json:"action_links"`
														} `json:"attachments"`
														SponsoredData interface{} `json:"sponsored_data"`
														Id            string      `json:"id"`
													} `json:"story"`
													HasCommerceAttachment                               bool `json:"has_commerce_attachment"`
													ModuleOperationCometFeedStoryActorPhotoSectionStory struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_CometFeedStoryActorPhotoSection_story"`
													ModuleComponentCometFeedStoryActorPhotoSectionStory struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_CometFeedStoryActorPhotoSection_story"`
												} `json:"actor_photo"`
												Metadata []struct {
													Typename             string      `json:"__typename"`
													IsICometStorySection string      `json:"__isICometStorySection"`
													IsProdEligible       bool        `json:"is_prod_eligible"`
													OverrideUrl          interface{} `json:"override_url"`
													VideoOverrideUrl     interface{} `json:"video_override_url"`
													Story                struct {
														CreationTime int    `json:"creation_time,omitempty"`
														Url          string `json:"url,omitempty"`
														GhlLabel     struct {
															Attributes []interface{} `json:"attributes"`
															Styles     []struct {
																Name string `json:"name"`
																Val  string `json:"val"`
															} `json:"styles"`
															Text     interface{} `json:"text"`
															Tag      string      `json:"tag"`
															Children []struct {
																Attributes []struct {
																	Name string `json:"name"`
																	Val  string `json:"val"`
																} `json:"attributes"`
																Styles []struct {
																	Name string `json:"name"`
																	Val  string `json:"val"`
																} `json:"styles"`
																Text     string        `json:"text"`
																Tag      string        `json:"tag"`
																Children []interface{} `json:"children"`
															} `json:"children"`
														} `json:"ghl_label,omitempty"`
														Id           string `json:"id"`
														PrivacyScope struct {
															IconImage struct {
																Name string `json:"name"`
															} `json:"icon_image"`
															Description string `json:"description"`
														} `json:"privacy_scope,omitempty"`
													} `json:"story"`
													ModuleOperationCometFeedStoryMetadataSectionStory struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_CometFeedStoryMetadataSection_story"`
													ModuleComponentCometFeedStoryMetadataSectionStory struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_CometFeedStoryMetadataSection_story"`
												} `json:"metadata"`
												Title struct {
													Typename             string `json:"__typename"`
													IsICometStorySection string `json:"__isICometStorySection"`
													IsProdEligible       bool   `json:"is_prod_eligible"`
													Story                struct {
														Id     string `json:"id"`
														Actors []struct {
															Typename              string      `json:"__typename"`
															Name                  string      `json:"name"`
															Id                    string      `json:"id"`
															IsActor               string      `json:"__isActor"`
															IsEntity              string      `json:"__isEntity"`
															Url                   string      `json:"url"`
															WorkForeignEntityInfo interface{} `json:"work_foreign_entity_info"`
															WorkInfo              interface{} `json:"work_info"`
														} `json:"actors"`
														Title         interface{} `json:"title"`
														CometSections struct {
															ActionLink   interface{} `json:"action_link"`
															Badge        interface{} `json:"badge"`
															FollowButton interface{} `json:"follow_button"`
														} `json:"comet_sections"`
														EncryptedTracking string `json:"encrypted_tracking"`
													} `json:"story"`
													ModuleOperationCometFeedStoryTitleSectionStory struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_CometFeedStoryTitleSection_story"`
													ModuleComponentCometFeedStoryTitleSectionStory struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_CometFeedStoryTitleSection_story"`
												} `json:"title"`
											} `json:"comet_sections"`
											EncryptedTracking   string      `json:"encrypted_tracking"`
											EasyHideButtonStory interface{} `json:"easy_hide_button_story"`
										} `json:"story"`
										IsRegulationEnforced                             bool `json:"is_regulation_enforced"`
										ModuleOperationCometFeedStoryContextSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryContextSection_story"`
										ModuleComponentCometFeedStoryContextSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryContextSection_story"`
									} `json:"context_layout"`
									AboveMessage      interface{} `json:"above_message"`
									InfoIcon          interface{} `json:"info_icon"`
									AttachmentOverlay interface{} `json:"attachment_overlay"`
									AttachedStory     interface{} `json:"attached_story"`
									Message           struct {
										Typename             string `json:"__typename"`
										IsICometStorySection string `json:"__isICometStorySection"`
										IsProdEligible       bool   `json:"is_prod_eligible"`
										Story                struct {
											IsTextOnlyStory bool `json:"is_text_only_story"`
											Message         struct {
												DelightRanges     []interface{} `json:"delight_ranges"`
												ImageRanges       []interface{} `json:"image_ranges"`
												InlineStyleRanges []interface{} `json:"inline_style_ranges"`
												AggregatedRanges  []interface{} `json:"aggregated_ranges"`
												Ranges            []struct {
													Entity struct {
														Typename  string `json:"__typename"`
														IsEntity  string `json:"__isEntity"`
														Url       string `json:"url"`
														MobileUrl string `json:"mobileUrl"`
														IsNode    string `json:"__isNode"`
														Id        string `json:"id"`
													} `json:"entity"`
													EntityIsWeakReference bool `json:"entity_is_weak_reference"`
													Length                int  `json:"length"`
													Offset                int  `json:"offset"`
												} `json:"ranges"`
												ColorRanges []interface{} `json:"color_ranges"`
												Text        string        `json:"text"`
											} `json:"message"`
											MessageTruncationLineLimit interface{} `json:"message_truncation_line_limit"`
											Id                         string      `json:"id"`
										} `json:"story"`
										ModuleOperationCometFeedStoryMessageSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryMessageSection_story"`
										ModuleComponentCometFeedStoryMessageSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryMessageSection_story"`
									} `json:"message"`
									MessageSuffix    interface{} `json:"message_suffix"`
									MessageContainer struct {
										Typename             string `json:"__typename"`
										IsICometStorySection string `json:"__isICometStorySection"`
										IsProdEligible       bool   `json:"is_prod_eligible"`
										Story                struct {
											Message struct {
												Text string `json:"text"`
											} `json:"message"`
											ReferencedSticker interface{} `json:"referenced_sticker"`
											Attachments       []struct {
												StyleList []string `json:"style_list"`
											} `json:"attachments"`
											TextFormatMetadata interface{} `json:"text_format_metadata"`
											CometSections      struct {
												Message interface{} `json:"message"`
											} `json:"comet_sections"`
											Id string `json:"id"`
										} `json:"story"`
										ModuleOperationCometFeedStoryMessageContainerSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryMessageContainerSection_story"`
										ModuleComponentCometFeedStoryMessageContainerSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryMessageContainerSection_story"`
									} `json:"message_container"`
									MessageSticker    interface{} `json:"message_sticker"`
									AggregatedStories interface{} `json:"aggregated_stories"`
								} `json:"comet_sections"`
								EncryptedTracking          string `json:"encrypted_tracking"`
								ShouldHostActorLinkInWatch bool   `json:"should_host_actor_link_in_watch"`
								Message                    struct {
									Text     string `json:"text"`
									Typename string `json:"__typename"`
								} `json:"message"`
								Attachments []struct {
									DeduplicationKey string `json:"deduplication_key"`
									Target           struct {
										Typename string `json:"__typename"`
										Id       string `json:"id"`
									} `json:"target"`
									Typename  string   `json:"__typename"`
									StyleList []string `json:"style_list"`
									Styles    struct {
										Typename                            string `json:"__typename"`
										IsStoryAttachmentStyleRendererUnion string `json:"__isStoryAttachmentStyleRendererUnion"`
										IsProdEligible                      bool   `json:"is_prod_eligible"`
										Attachment                          struct {
											Url               string        `json:"url"`
											ActionLinks       []interface{} `json:"action_links"`
											CtaScreenRenderer interface{}   `json:"cta_screen_renderer"`
											Media             struct {
												Typename       string `json:"__typename"`
												ThumbnailImage struct {
													Uri string `json:"uri"`
												} `json:"thumbnailImage"`
												Id                string `json:"id"`
												IsClippingEnabled bool   `json:"is_clipping_enabled"`
												LiveRewindEnabled bool   `json:"live_rewind_enabled"`
												Owner             struct {
													Typename                        string `json:"__typename"`
													Id                              string `json:"id"`
													IsVideoOwner                    string `json:"__isVideoOwner"`
													HasProfessionalFeaturesForWatch bool   `json:"has_professional_features_for_watch"`
												} `json:"owner"`
												PlayableDurationInMs                   int           `json:"playable_duration_in_ms"`
												IsHuddle                               bool          `json:"is_huddle"`
												Url                                    string        `json:"url"`
												IfViewerCanUseLatencyMenu              interface{}   `json:"if_viewer_can_use_latency_menu"`
												IfViewerCanUseLatencyMenuToggle        interface{}   `json:"if_viewer_can_use_latency_menu_toggle"`
												CaptionsUrl                            interface{}   `json:"captions_url"`
												VideoAvailableCaptionsLocales          []interface{} `json:"video_available_captions_locales"`
												IfViewerCanSeeCommunityModerationTools interface{}   `json:"if_viewer_can_see_community_moderation_tools"`
												IfViewerCanUseLiveRewind               interface{}   `json:"if_viewer_can_use_live_rewind"`
												IfViewerCanUseClipping                 interface{}   `json:"if_viewer_can_use_clipping"`
												IfViewerCanSeeCostreamingTools         interface{}   `json:"if_viewer_can_see_costreaming_tools"`
												VideoPlayerScrubberPreviewRenderer     struct {
													Typename string `json:"__typename"`
													Video    struct {
														ScrubberPreviewThumbnailInformation struct {
															SpriteUris                 []string `json:"sprite_uris"`
															ThumbnailWidth             int      `json:"thumbnail_width"`
															ThumbnailHeight            int      `json:"thumbnail_height"`
															HasPreviewThumbnails       bool     `json:"has_preview_thumbnails"`
															NumImagesPerRow            int      `json:"num_images_per_row"`
															MaxNumberOfImagesPerSprite int      `json:"max_number_of_images_per_sprite"`
															TimeIntervalBetweenImage   int      `json:"time_interval_between_image"`
														} `json:"scrubber_preview_thumbnail_information"`
														Id string `json:"id"`
													} `json:"video"`
													ModuleOperationVideoPlayerScrubberPreviewVideo struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_VideoPlayerScrubberPreview_video"`
													ModuleComponentVideoPlayerScrubberPreviewVideo struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_VideoPlayerScrubberPreview_video"`
												} `json:"video_player_scrubber_preview_renderer"`
												VideoPlayerScrubberBaseContentRenderer interface{} `json:"video_player_scrubber_base_content_renderer"`
												RecipientGroup                         interface{} `json:"recipient_group"`
												MusicAttachmentMetadata                interface{} `json:"music_attachment_metadata"`
												VideoContainerType                     string      `json:"video_container_type"`
												BreakingStatus                         bool        `json:"breakingStatus"`
												VideoId                                string      `json:"videoId"`
												IsPremiere                             bool        `json:"isPremiere"`
												LiveViewerCount                        int         `json:"liveViewerCount"`
												RehearsalInfo                          interface{} `json:"rehearsalInfo"`
												IsGamingVideo                          bool        `json:"is_gaming_video"`
												IsLiveAudioRoomV2Broadcast             bool        `json:"is_live_audio_room_v2_broadcast"`
												PublishTime                            int         `json:"publish_time"`
												LiveSpeakerCountIndicator              interface{} `json:"live_speaker_count_indicator"`
												CanViewerShare                         bool        `json:"can_viewer_share"`
												CreationStory                          struct {
													Shareable struct {
														Typename string `json:"__typename"`
														WwwUrl   string `json:"wwwUrl"`
														IsNode   string `json:"__isNode"`
														Id       string `json:"id"`
													} `json:"shareable"`
													Id string `json:"id"`
												} `json:"creation_story"`
												EndCardsChannelInfo                    interface{} `json:"end_cards_channel_info"`
												IsSoundbitesVideo                      bool        `json:"is_soundbites_video"`
												IsLooping                              bool        `json:"is_looping"`
												Info                                   interface{} `json:"info"`
												AnimatedImageCaption                   interface{} `json:"animated_image_caption"`
												Width                                  int         `json:"width"`
												Height                                 int         `json:"height"`
												BroadcasterOrigin                      interface{} `json:"broadcaster_origin"`
												BroadcastId                            interface{} `json:"broadcast_id"`
												BroadcastStatus                        interface{} `json:"broadcast_status"`
												IsLiveStreaming                        bool        `json:"is_live_streaming"`
												IsLiveTraceEnabled                     bool        `json:"is_live_trace_enabled"`
												IsVideoBroadcast                       bool        `json:"is_video_broadcast"`
												IsPodcastVideo                         bool        `json:"is_podcast_video"`
												LoopCount                              int         `json:"loop_count"`
												IsSpherical                            bool        `json:"is_spherical"`
												IsSphericalEnabled                     bool        `json:"is_spherical_enabled"`
												UnsupportedBrowserMessage              interface{} `json:"unsupported_browser_message"`
												PmvMetadata                            interface{} `json:"pmv_metadata"`
												LatencySensitiveConfig                 interface{} `json:"latency_sensitive_config"`
												LivePlaybackInstrumentationConfigs     interface{} `json:"live_playback_instrumentation_configs"`
												IsNcsr                                 bool        `json:"is_ncsr"`
												PermalinkUrl                           string      `json:"permalink_url"`
												DashPrefetchExperimental               []string    `json:"dash_prefetch_experimental"`
												VideoStatusType                        string      `json:"video_status_type"`
												CanUseOz                               bool        `json:"can_use_oz"`
												DashManifest                           string      `json:"dash_manifest"`
												DashManifestUrl                        string      `json:"dash_manifest_url"`
												MinQualityPreference                   interface{} `json:"min_quality_preference"`
												AudioUserPreferredLanguage             string      `json:"audio_user_preferred_language"`
												IsRssPodcastVideo                      bool        `json:"is_rss_podcast_video"`
												BrowserNativeSdUrl                     string      `json:"browser_native_sd_url"`
												BrowserNativeHdUrl                     string      `json:"browser_native_hd_url"`
												SphericalVideoFallbackUrls             interface{} `json:"spherical_video_fallback_urls"`
												IsLatencyMenuEnabled                   bool        `json:"is_latency_menu_enabled"`
												FblsTier                               interface{} `json:"fbls_tier"`
												IsLatencySensitiveBroadcast            bool        `json:"is_latency_sensitive_broadcast"`
												CometVideoPlayerStaticConfig           string      `json:"comet_video_player_static_config"`
												CometVideoPlayerContextSensitiveConfig string      `json:"comet_video_player_context_sensitive_config"`
												VideoPlayerShakaPerformanceLoggerInit  struct {
													Typename                                                          string `json:"__typename"`
													ModuleOperationUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
													ModuleComponentUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
												} `json:"video_player_shaka_performance_logger_init"`
												VideoPlayerShakaPerformanceLoggerShouldSample bool `json:"video_player_shaka_performance_logger_should_sample"`
												VideoPlayerShakaPerformanceLoggerInit2        struct {
													Typename                                                        string `json:"__typename"`
													ModuleOperationUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
													ModuleComponentUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
													PerSessionSamplingRate interface{} `json:"per_session_sampling_rate"`
												} `json:"video_player_shaka_performance_logger_init2"`
												AutoplayGatingResult      string        `json:"autoplay_gating_result"`
												ViewerAutoplaySetting     string        `json:"viewer_autoplay_setting"`
												CanAutoplay               bool          `json:"can_autoplay"`
												DrmInfo                   string        `json:"drm_info"`
												P2PSettings               interface{}   `json:"p2p_settings"`
												AudioSettings             interface{}   `json:"audio_settings"`
												CaptionsSettings          interface{}   `json:"captions_settings"`
												BroadcastLowLatencyConfig interface{}   `json:"broadcast_low_latency_config"`
												AudioAvailability         string        `json:"audio_availability"`
												MutedSegments             []interface{} `json:"muted_segments"`
												SphericalVideoRenderer    interface{}   `json:"spherical_video_renderer"`
												PreferredThumbnail        struct {
													Image struct {
														Uri string `json:"uri"`
													} `json:"image"`
													ImagePreviewPayload string `json:"image_preview_payload"`
													Id                  string `json:"id"`
												} `json:"preferred_thumbnail"`
												VideoImfData                                  interface{} `json:"video_imf_data"`
												OriginalWidth                                 int         `json:"original_width"`
												OriginalHeight                                int         `json:"original_height"`
												OriginalRotation                              string      `json:"original_rotation"`
												IfViewerCanSeePayToAccessPaywall              interface{} `json:"if_viewer_can_see_pay_to_access_paywall"`
												CometVideoPlayerAudioOverlayRenderer          interface{} `json:"comet_video_player_audio_overlay_renderer"`
												CometVideoPlayerAudioBackgroundRenderer       interface{} `json:"comet_video_player_audio_background_renderer"`
												CometVideoPlayerMusicSproutBackgroundRenderer interface{} `json:"comet_video_player_music_sprout_background_renderer"`
												CometProductTagFeedOverlayRenderer            interface{} `json:"comet_product_tag_feed_overlay_renderer"`
												ClipFallbackCover                             interface{} `json:"clip_fallback_cover"`
												IsClip                                        bool        `json:"is_clip"`
												MatchaRelatedKeywordsLinks                    []string    `json:"matcha_related_keywords_links"`
												IsMusicClip                                   bool        `json:"is_music_clip"`
												VideoCollaboratorPageOrDelegatePage           interface{} `json:"video_collaborator_page_or_delegate_page"`
												VideoAnchorTagInfo                            interface{} `json:"video_anchor_tag_info"`
												Image                                         struct {
													Uri string `json:"uri"`
												} `json:"image"`
												CanonicalUriWithFallback string `json:"canonical_uri_with_fallback"`
												IsNode                   string `json:"__isNode"`
											} `json:"media"`
										} `json:"attachment"`
										ModuleOperationCometFeedStoryAttachmentRendererInnerAttachment struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryAttachmentRendererInner_attachment"`
										ModuleComponentCometFeedStoryAttachmentRendererInnerAttachment struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryAttachmentRendererInner_attachment"`
									} `json:"styles"`
									ThrowbackStyles     interface{} `json:"throwbackStyles"`
									CometFooterRenderer struct {
										Typename   string `json:"__typename"`
										Attachment struct {
											GhlMockedFooterInfo struct {
												Headline   string `json:"headline"`
												FooterBody string `json:"footer_body"`
												Link       string `json:"link"`
												Meta       string `json:"meta"`
												CtaButton  struct {
													Attributes []interface{} `json:"attributes"`
													Styles     []struct {
														Name string `json:"name"`
														Val  string `json:"val"`
													} `json:"styles"`
													Text     interface{} `json:"text"`
													Tag      string      `json:"tag"`
													Children []struct {
														Attributes []struct {
															Name string `json:"name"`
															Val  string `json:"val"`
														} `json:"attributes"`
														Styles []struct {
															Name string `json:"name"`
															Val  string `json:"val"`
														} `json:"styles"`
														Text     string        `json:"text"`
														Tag      string        `json:"tag"`
														Children []interface{} `json:"children"`
													} `json:"children"`
												} `json:"cta_button"`
											} `json:"ghl_mocked_footer_info"`
										} `json:"attachment"`
										ModuleOperationCometFeedStoryAttachmentFooterSectionAttachment struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryAttachmentFooterSection_attachment"`
										ModuleComponentCometFeedStoryAttachmentFooterSectionAttachment struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryAttachmentFooterSection_attachment"`
									} `json:"comet_footer_renderer"`
									CometFooterDisclaimerRenderer interface{} `json:"comet_footer_disclaimer_renderer"`
								} `json:"attachments"`
								SponsoredData      interface{} `json:"sponsored_data"`
								TextFormatMetadata interface{} `json:"text_format_metadata"`
								Actors             []struct {
									Typename string `json:"__typename"`
									Id       string `json:"id"`
									Name     string `json:"name"`
									IsEntity string `json:"__isEntity"`
									Url      string `json:"url"`
								} `json:"actors"`
								GhlMockedEncryptedLink  string      `json:"ghl_mocked_encrypted_link"`
								GhlLabelMockedCtaButton interface{} `json:"ghl_label_mocked_cta_button"`
								WwwURL                  string      `json:"wwwURL"`
								TargetGroup             interface{} `json:"target_group"`
								AttachedStory           interface{} `json:"attached_story"`
							} `json:"story"`
							ModuleOperationCometFeedStoryStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_operation_CometFeedStory_story"`
							ModuleComponentCometFeedStoryStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_component_CometFeedStory_story"`
						} `json:"content"`
						Layout struct {
							Typename                                 string `json:"__typename"`
							IsICometStorySection                     string `json:"__isICometStorySection"`
							IsProdEligible                           bool   `json:"is_prod_eligible"`
							ModuleOperationCometFeedStoryStoryLayout struct {
								Dr string `json:"__dr"`
							} `json:"__module_operation_CometFeedStory_story__layout"`
							ModuleComponentCometFeedStoryStoryLayout struct {
								Dr string `json:"__dr"`
							} `json:"__module_component_CometFeedStory_story__layout"`
						} `json:"layout"`
						CopyrightViolationHeader interface{} `json:"copyright_violation_header"`
						Header                   interface{} `json:"header"`
						ContextLayout            struct {
							Typename                      string      `json:"__typename"`
							IsICometStorySection          string      `json:"__isICometStorySection"`
							IsProdEligible                bool        `json:"is_prod_eligible"`
							LocalAlertsStoryMenuPromotion interface{} `json:"local_alerts_story_menu_promotion"`
							Story                         struct {
								Id                        string      `json:"id"`
								DebugInfo                 interface{} `json:"debug_info"`
								SerializedFrtpIdentifiers interface{} `json:"serialized_frtp_identifiers"`
								CanViewerSeeMenu          bool        `json:"can_viewer_see_menu"`
								CometSections             struct {
									ActorPhoto struct {
										Typename             string `json:"__typename"`
										IsICometStorySection string `json:"__isICometStorySection"`
										IsProdEligible       bool   `json:"is_prod_eligible"`
										Story                struct {
											Actors []struct {
												Typename              string      `json:"__typename"`
												IsActor               string      `json:"__isActor"`
												Id                    string      `json:"id"`
												IsEntity              string      `json:"__isEntity"`
												Url                   string      `json:"url"`
												WorkForeignEntityInfo interface{} `json:"work_foreign_entity_info"`
												WorkInfo              interface{} `json:"work_info"`
												StoryBucket           struct {
													Nodes []struct {
														ShouldShowCloseFriendBadge bool   `json:"should_show_close_friend_badge"`
														Id                         string `json:"id"`
														FirstStoryToShow           struct {
															Id                 string `json:"id"`
															StoryCardSeenState struct {
																IsSeenByViewer bool `json:"is_seen_by_viewer"`
															} `json:"story_card_seen_state"`
														} `json:"first_story_to_show"`
													} `json:"nodes"`
												} `json:"story_bucket"`
												LiveVideoForCometLiveRing interface{} `json:"live_video_for_comet_live_ring"`
												AnswerAgentGroupId        interface{} `json:"answer_agent_group_id"`
												ProfileUrl                string      `json:"profile_url"`
												Name                      string      `json:"name"`
												ProfilePicture            struct {
													Uri    string `json:"uri"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
													Scale  int    `json:"scale"`
												} `json:"profile_picture"`
												IsAdditionalProfilePlus bool        `json:"is_additional_profile_plus"`
												DelegatePage            interface{} `json:"delegate_page"`
											} `json:"actors"`
											CometSections struct {
												ActionLink interface{} `json:"action_link"`
											} `json:"comet_sections"`
											Attachments []struct {
												ActionLinks []interface{} `json:"action_links"`
											} `json:"attachments"`
											SponsoredData interface{} `json:"sponsored_data"`
											Id            string      `json:"id"`
										} `json:"story"`
										HasCommerceAttachment                               bool `json:"has_commerce_attachment"`
										ModuleOperationCometFeedStoryActorPhotoSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryActorPhotoSection_story"`
										ModuleComponentCometFeedStoryActorPhotoSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryActorPhotoSection_story"`
									} `json:"actor_photo"`
									Metadata []struct {
										Typename             string      `json:"__typename"`
										IsICometStorySection string      `json:"__isICometStorySection"`
										IsProdEligible       bool        `json:"is_prod_eligible"`
										OverrideUrl          interface{} `json:"override_url"`
										VideoOverrideUrl     interface{} `json:"video_override_url"`
										Story                struct {
											CreationTime int    `json:"creation_time,omitempty"`
											Url          string `json:"url,omitempty"`
											GhlLabel     struct {
												Attributes []interface{} `json:"attributes"`
												Styles     []struct {
													Name string `json:"name"`
													Val  string `json:"val"`
												} `json:"styles"`
												Text     interface{} `json:"text"`
												Tag      string      `json:"tag"`
												Children []struct {
													Attributes []struct {
														Name string `json:"name"`
														Val  string `json:"val"`
													} `json:"attributes"`
													Styles []struct {
														Name string `json:"name"`
														Val  string `json:"val"`
													} `json:"styles"`
													Text     string        `json:"text"`
													Tag      string        `json:"tag"`
													Children []interface{} `json:"children"`
												} `json:"children"`
											} `json:"ghl_label,omitempty"`
											Id           string `json:"id"`
											PrivacyScope struct {
												IconImage struct {
													Name string `json:"name"`
												} `json:"icon_image"`
												Description string `json:"description"`
											} `json:"privacy_scope,omitempty"`
										} `json:"story"`
										ModuleOperationCometFeedStoryMetadataSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryMetadataSection_story"`
										ModuleComponentCometFeedStoryMetadataSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryMetadataSection_story"`
									} `json:"metadata"`
									Title struct {
										Typename             string `json:"__typename"`
										IsICometStorySection string `json:"__isICometStorySection"`
										IsProdEligible       bool   `json:"is_prod_eligible"`
										Story                struct {
											Id     string `json:"id"`
											Actors []struct {
												Typename              string      `json:"__typename"`
												Name                  string      `json:"name"`
												Id                    string      `json:"id"`
												IsActor               string      `json:"__isActor"`
												IsEntity              string      `json:"__isEntity"`
												Url                   string      `json:"url"`
												WorkForeignEntityInfo interface{} `json:"work_foreign_entity_info"`
												WorkInfo              interface{} `json:"work_info"`
											} `json:"actors"`
											Title         interface{} `json:"title"`
											CometSections struct {
												ActionLink   interface{} `json:"action_link"`
												Badge        interface{} `json:"badge"`
												FollowButton interface{} `json:"follow_button"`
											} `json:"comet_sections"`
											EncryptedTracking string `json:"encrypted_tracking"`
										} `json:"story"`
										ModuleOperationCometFeedStoryTitleSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_operation_CometFeedStoryTitleSection_story"`
										ModuleComponentCometFeedStoryTitleSectionStory struct {
											Dr string `json:"__dr"`
										} `json:"__module_component_CometFeedStoryTitleSection_story"`
									} `json:"title"`
								} `json:"comet_sections"`
								EncryptedTracking   string      `json:"encrypted_tracking"`
								EasyHideButtonStory interface{} `json:"easy_hide_button_story"`
							} `json:"story"`
							IsRegulationEnforced                             bool `json:"is_regulation_enforced"`
							ModuleOperationCometFeedStoryContextSectionStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_operation_CometFeedStoryContextSection_story"`
							ModuleComponentCometFeedStoryContextSectionStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_component_CometFeedStoryContextSection_story"`
						} `json:"context_layout"`
						AymtFooter interface{} `json:"aymt_footer"`
						Footer     interface{} `json:"footer"`
						Feedback   struct {
							Typename             string `json:"__typename"`
							IsICometStorySection string `json:"__isICometStorySection"`
							IsProdEligible       bool   `json:"is_prod_eligible"`
							Story                struct {
								IsTextOnlyStory bool `json:"is_text_only_story"`
								FeedbackContext struct {
									FeedbackTargetWithContext struct {
										ViewerActor struct {
											Typename string `json:"__typename"`
											Id       string `json:"id"`
										} `json:"viewer_actor"`
										UfiRenderer struct {
											Typename string `json:"__typename"`
											Feedback struct {
												IsEligibleForRealTimeUpdates        bool        `json:"is_eligible_for_real_time_updates"`
												CanViewerComment                    bool        `json:"can_viewer_comment"`
												AssociatedGroup                     interface{} `json:"associated_group"`
												RecruitingGroupUnencodedId          interface{} `json:"recruiting_group_unencoded_id"`
												Id                                  string      `json:"id"`
												IsEligibleForEnhancedCommentUpdates bool        `json:"is_eligible_for_enhanced_comment_updates"`
												IsSyncedQaPost                      bool        `json:"is_synced_qa_post"`
												SubscriptionTargetId                string      `json:"subscription_target_id"`
												AssociatedVideo                     struct {
													BroadcastIsAmaEnabled bool   `json:"broadcast_is_ama_enabled"`
													Id                    string `json:"id"`
												} `json:"associated_video"`
												CommentListRenderer struct {
													Typename string `json:"__typename"`
													Feedback struct {
														TotalCommentCount                       int `json:"total_comment_count"`
														CommentRenderingInstanceForFeedLocation struct {
															SelectedIntent struct {
																Title       string `json:"title"`
																Description string `json:"description"`
																IntentToken string `json:"intent_token"`
															} `json:"selected_intent"`
															SelectableIntents []struct {
																Title       string `json:"title"`
																Description string `json:"description"`
																IntentToken string `json:"intent_token"`
															} `json:"selectable_intents"`
															Comments struct {
																CreatedCommentInsertionPosition string `json:"created_comment_insertion_position"`
																FilteringFooterString           string `json:"filtering_footer_string"`
																Count                           int    `json:"count"`
																PageSize                        int    `json:"page_size"`
																TotalCount                      int    `json:"total_count"`
																Edges                           []struct {
																	Node struct {
																		Id       string `json:"id"`
																		Feedback struct {
																			Id            string `json:"id"`
																			ExpansionInfo struct {
																				ExpansionToken       string `json:"expansion_token"`
																				ShouldShowReplyCount bool   `json:"should_show_reply_count"`
																			} `json:"expansion_info"`
																			RepliesFields struct {
																				FilteringFooterString string `json:"filtering_footer_string"`
																				Count                 int    `json:"count"`
																				TotalCount            int    `json:"total_count"`
																			} `json:"replies_fields"`
																			ViewerActor struct {
																				Typename             string `json:"__typename"`
																				Id                   string `json:"id"`
																				IsActor              string `json:"__isActor"`
																				Name                 string `json:"name"`
																				ProfilePictureDepth0 struct {
																					Uri string `json:"uri"`
																				} `json:"profile_picture_depth_0"`
																				ProfilePictureDepth1 struct {
																					Uri string `json:"uri"`
																				} `json:"profile_picture_depth_1"`
																				Gender string `json:"gender"`
																			} `json:"viewer_actor"`
																			Url      string `json:"url"`
																			Typename string `json:"__typename"`
																			Plugins  []struct {
																				Typename                                          string      `json:"__typename"`
																				ContextId                                         interface{} `json:"context_id"`
																				PostId                                            string      `json:"post_id,omitempty"`
																				ModuleOperationUseCometUFIComposerPluginsFeedback struct {
																					Dr string `json:"__dr"`
																				} `json:"__module_operation_useCometUFIComposerPlugins_feedback"`
																				ModuleComponentUseCometUFIComposerPluginsFeedback struct {
																					Dr string `json:"__dr"`
																				} `json:"__module_component_useCometUFIComposerPlugins_feedback"`
																				HasAvatar   bool   `json:"has_avatar,omitempty"`
																				FeedbackId  string `json:"feedback_id,omitempty"`
																				EmojiSize   int    `json:"emoji_size,omitempty"`
																				ViewerActor struct {
																					Typename string `json:"__typename"`
																					Id       string `json:"id"`
																				} `json:"viewer_actor,omitempty"`
																				ShouldCondenseVideoPreview bool   `json:"should_condense_video_preview,omitempty"`
																				OwningProfileId            string `json:"owning_profile_id,omitempty"`
																			} `json:"plugins"`
																			CommentComposerPlaceholder     string      `json:"comment_composer_placeholder"`
																			CanViewerComment               bool        `json:"can_viewer_comment"`
																			ConstituentBadgeBannerRenderer interface{} `json:"constituent_badge_banner_renderer"`
																			HaveCommentsBeenDisabled       bool        `json:"have_comments_been_disabled"`
																			AreLiveVideoCommentsDisabled   bool        `json:"are_live_video_comments_disabled"`
																			IsViewerMuted                  bool        `json:"is_viewer_muted"`
																			TotalCommentCount              int         `json:"total_comment_count"`
																			CommentsDisabledNoticeRenderer struct {
																				Typename      string `json:"__typename"`
																				NoticeMessage struct {
																					DelightRanges     []interface{} `json:"delight_ranges"`
																					ImageRanges       []interface{} `json:"image_ranges"`
																					InlineStyleRanges []interface{} `json:"inline_style_ranges"`
																					AggregatedRanges  []interface{} `json:"aggregated_ranges"`
																					Ranges            []interface{} `json:"ranges"`
																					ColorRanges       []interface{} `json:"color_ranges"`
																					Text              string        `json:"text"`
																				} `json:"notice_message"`
																				ModuleOperationCometUFICommentDisabledNoticeFeedback struct {
																					Dr string `json:"__dr"`
																				} `json:"__module_operation_CometUFICommentDisabledNotice_feedback"`
																				ModuleComponentCometUFICommentDisabledNoticeFeedback struct {
																					Dr string `json:"__dr"`
																				} `json:"__module_component_CometUFICommentDisabledNotice_feedback"`
																			} `json:"comments_disabled_notice_renderer"`
																			RepliesConnection struct {
																				Edges    []interface{} `json:"edges"`
																				PageInfo struct {
																					EndCursor       interface{} `json:"end_cursor"`
																					HasNextPage     bool        `json:"has_next_page"`
																					HasPreviousPage bool        `json:"has_previous_page"`
																					StartCursor     interface{} `json:"start_cursor"`
																				} `json:"page_info"`
																			} `json:"replies_connection"`
																			ParentObjectEnt struct {
																				Typename                      string      `json:"__typename"`
																				InlineRepliesExpanderRenderer interface{} `json:"inline_replies_expander_renderer"`
																				Id                            string      `json:"id"`
																			} `json:"parent_object_ent"`
																			ViewerFeedbackReactionInfo interface{} `json:"viewer_feedback_reaction_info"`
																			TopReactions               struct {
																				Edges []interface{} `json:"edges"`
																			} `json:"top_reactions"`
																			Reactors struct {
																				CountReduced string `json:"count_reduced"`
																				Count        int    `json:"count"`
																			} `json:"reactors"`
																			CanSeeTopCustomReactionsOnCommentForCreators interface{} `json:"can_see_top_custom_reactions_on_comment_for_creators"`
																			ShouldShowTopReactions                       bool        `json:"should_show_top_reactions"`
																		} `json:"feedback"`
																		LegacyFbid string `json:"legacy_fbid"`
																		Depth      int    `json:"depth"`
																		Body       struct {
																			Text   string        `json:"text"`
																			Ranges []interface{} `json:"ranges"`
																		} `json:"body"`
																		Attachments                    []interface{} `json:"attachments"`
																		IsMarkdownEnabled              bool          `json:"is_markdown_enabled"`
																		CommunityCommentSignalRenderer interface{}   `json:"community_comment_signal_renderer"`
																		CommentMenuTooltip             string        `json:"comment_menu_tooltip"`
																		ShouldShowCommentMenu          bool          `json:"should_show_comment_menu"`
																		Author                         struct {
																			Typename             string `json:"__typename"`
																			Id                   string `json:"id"`
																			Name                 string `json:"name"`
																			IsActor              string `json:"__isActor"`
																			ProfilePictureDepth0 struct {
																				Uri string `json:"uri"`
																			} `json:"profile_picture_depth_0"`
																			ProfilePictureDepth1 struct {
																				Uri string `json:"uri"`
																			} `json:"profile_picture_depth_1"`
																			Gender          string      `json:"gender"`
																			IsEntity        string      `json:"__isEntity"`
																			Url             string      `json:"url"`
																			WorkInfo        interface{} `json:"work_info"`
																			IsVerified      bool        `json:"is_verified"`
																			ShortName       string      `json:"short_name"`
																			SubscribeStatus string      `json:"subscribe_status"`
																		} `json:"author"`
																		IsAuthorWeakReference bool `json:"is_author_weak_reference"`
																		CommentActionLinks    []struct {
																			Typename string `json:"__typename"`
																			Comment  struct {
																				Id                 string `json:"id"`
																				CreatedTime        int    `json:"created_time,omitempty"`
																				Url                string `json:"url,omitempty"`
																				IsLiveVideoComment bool   `json:"is_live_video_comment,omitempty"`
																				Feedback           struct {
																					Id                         string      `json:"id"`
																					ViewerFeedbackReactionInfo interface{} `json:"viewer_feedback_reaction_info"`
																					ViewerActor                struct {
																						Typename string `json:"__typename"`
																						Id       string `json:"id"`
																					} `json:"viewer_actor"`
																					SupportedReactionInfos []struct {
																						Animation struct {
																							UriKeyframes2 string `json:"uri_keyframes2"`
																						} `json:"animation"`
																						Id string `json:"id"`
																					} `json:"supported_reaction_infos"`
																					IfViewerCanUseCreatorCustomReactions interface{} `json:"if_viewer_can_use_creator_custom_reactions"`
																					AssociatedVideo                      interface{} `json:"associated_video"`
																					TopReactions                         struct {
																						Edges []interface{} `json:"edges"`
																					} `json:"top_reactions"`
																					UnifiedReactors struct {
																						Count int `json:"count"`
																					} `json:"unified_reactors"`
																					Reactors struct {
																						Count   int  `json:"count"`
																						IsEmpty bool `json:"is_empty"`
																					} `json:"reactors"`
																					IfViewerCanRenderCreatorCustomReactions interface{} `json:"if_viewer_can_render_creator_custom_reactions"`
																				} `json:"feedback,omitempty"`
																				PreferredBody struct {
																					Typename        string `json:"__typename"`
																					TranslationType string `json:"translation_type"`
																				} `json:"preferred_body,omitempty"`
																				TranslatabilityForViewer struct {
																					SourceDialectName string `json:"source_dialect_name"`
																				} `json:"translatability_for_viewer,omitempty"`
																			} `json:"comment"`
																			ModuleOperationCometUFICommentActionLinksComment struct {
																				Dr string `json:"__dr"`
																			} `json:"__module_operation_CometUFICommentActionLinks_comment"`
																			ModuleComponentCometUFICommentActionLinksComment struct {
																				Dr string `json:"__dr"`
																			} `json:"__module_component_CometUFICommentActionLinks_comment"`
																		} `json:"comment_action_links"`
																		PreferredBody struct {
																			Typename          string        `json:"__typename"`
																			DelightRanges     []interface{} `json:"delight_ranges"`
																			ImageRanges       []interface{} `json:"image_ranges"`
																			InlineStyleRanges []interface{} `json:"inline_style_ranges"`
																			AggregatedRanges  []interface{} `json:"aggregated_ranges"`
																			Ranges            []interface{} `json:"ranges"`
																			ColorRanges       []interface{} `json:"color_ranges"`
																			Text              string        `json:"text"`
																			TranslationType   string        `json:"translation_type"`
																		} `json:"preferred_body"`
																		BodyRenderer struct {
																			Typename                                              string        `json:"__typename"`
																			DelightRanges                                         []interface{} `json:"delight_ranges"`
																			ImageRanges                                           []interface{} `json:"image_ranges"`
																			InlineStyleRanges                                     []interface{} `json:"inline_style_ranges"`
																			AggregatedRanges                                      []interface{} `json:"aggregated_ranges"`
																			Ranges                                                []interface{} `json:"ranges"`
																			ColorRanges                                           []interface{} `json:"color_ranges"`
																			Text                                                  string        `json:"text"`
																			ModuleOperationCometUFICommentTextBodyRendererComment struct {
																				Dr string `json:"__dr"`
																			} `json:"__module_operation_CometUFICommentTextBodyRenderer_comment"`
																			ModuleComponentCometUFICommentTextBodyRendererComment struct {
																				Dr string `json:"__dr"`
																			} `json:"__module_component_CometUFICommentTextBodyRenderer_comment"`
																		} `json:"body_renderer"`
																		CommentParent                   interface{} `json:"comment_parent"`
																		IsDeclinedByGroupAdminAssistant bool        `json:"is_declined_by_group_admin_assistant"`
																		IsGamingVideoComment            bool        `json:"is_gaming_video_comment"`
																		TimestampInVideo                int         `json:"timestamp_in_video"`
																		TranslatabilityForViewer        struct {
																			SourceDialect string `json:"source_dialect"`
																		} `json:"translatability_for_viewer"`
																		WrittenWhileVideoWasLive     bool        `json:"written_while_video_was_live"`
																		GroupCommentInfo             interface{} `json:"group_comment_info"`
																		BizwebCommentInfo            interface{} `json:"bizweb_comment_info"`
																		HasConstituentBadge          bool        `json:"has_constituent_badge"`
																		CanViewerSeeSubsribeButton   bool        `json:"can_viewer_see_subsribe_button"`
																		CanSeeConstituentBadgeUpsell bool        `json:"can_see_constituent_badge_upsell"`
																		LegacyToken                  string      `json:"legacy_token"`
																		ParentFeedback               struct {
																			Id                  string      `json:"id"`
																			ShareFbid           string      `json:"share_fbid"`
																			PoliticalFigureData interface{} `json:"political_figure_data"`
																			OwningProfile       struct {
																				Typename string `json:"__typename"`
																				Name     string `json:"name"`
																				Id       string `json:"id"`
																			} `json:"owning_profile"`
																		} `json:"parent_feedback"`
																		QuestionAndAnswerType         interface{}   `json:"question_and_answer_type"`
																		IsAuthorOriginalPoster        bool          `json:"is_author_original_poster"`
																		IsViewerCommentPoster         bool          `json:"is_viewer_comment_poster"`
																		IsAuthorBot                   bool          `json:"is_author_bot"`
																		IsAuthorNonCoworker           bool          `json:"is_author_non_coworker"`
																		AuthorUserSignalsRenderer     interface{}   `json:"author_user_signals_renderer"`
																		AuthorBadgeRenderers          []interface{} `json:"author_badge_renderers"`
																		IdentityBadgesWeb             []interface{} `json:"identity_badges_web"`
																		CanShowMultipleIdentityBadges bool          `json:"can_show_multiple_identity_badges"`
																		DiscoverableIdentityBadgesWeb []interface{} `json:"discoverable_identity_badges_web"`
																		User                          struct {
																			Name           string `json:"name"`
																			ProfilePicture struct {
																				Uri string `json:"uri"`
																			} `json:"profile_picture"`
																			Id string `json:"id"`
																		} `json:"user"`
																		ParentPostStory struct {
																			Attachments []struct {
																				Media struct {
																					Typename          string      `json:"__typename"`
																					Id                string      `json:"id"`
																					IsLiveStreaming   bool        `json:"is_live_streaming"`
																					BroadcastDuration interface{} `json:"broadcast_duration"`
																					PlayableDuration  int         `json:"playable_duration"`
																					Owner             struct {
																						Typename      string `json:"__typename"`
																						IsVideoOwner  string `json:"__isVideoOwner"`
																						ActorToFollow struct {
																							Typename                      string `json:"__typename"`
																							Id                            string `json:"id"`
																							VideoChannelIsViewerFollowing bool   `json:"video_channel_is_viewer_following"`
																						} `json:"actor_to_follow"`
																						Name string `json:"name"`
																						Id   string `json:"id"`
																					} `json:"owner"`
																					IsNode string `json:"__isNode"`
																				} `json:"media"`
																			} `json:"attachments"`
																			Id string `json:"id"`
																		} `json:"parent_post_story"`
																		WorkAmaAnswerStatus                               interface{}   `json:"work_ama_answer_status"`
																		WorkKnowledgeInlineAnnotationCommentBadgeRenderer interface{}   `json:"work_knowledge_inline_annotation_comment_badge_renderer"`
																		BusinessCommentAttributes                         []interface{} `json:"business_comment_attributes"`
																		IsLiveVideoComment                                bool          `json:"is_live_video_comment"`
																		CreatedTime                                       int           `json:"created_time"`
																		TranslationAvailableForViewer                     bool          `json:"translation_available_for_viewer"`
																		InlineSurveyConfig                                interface{}   `json:"inline_survey_config"`
																		SpamDisplayMode                                   string        `json:"spam_display_mode"`
																		AttachedStory                                     interface{}   `json:"attached_story"`
																		HelpfulCommentsCommentEyebrowRenderer             interface{}   `json:"helpful_comments_comment_eyebrow_renderer"`
																		CommentDirectParent                               interface{}   `json:"comment_direct_parent"`
																		IfViewerCanSeeMemberPageTooltip                   interface{}   `json:"if_viewer_can_see_member_page_tooltip"`
																		IsDisabled                                        bool          `json:"is_disabled"`
																		WorkAnsweredEventCommentRenderer                  interface{}   `json:"work_answered_event_comment_renderer"`
																		CommentUpperBadgeRenderer                         interface{}   `json:"comment_upper_badge_renderer"`
																		ElevatedCommentData                               interface{}   `json:"elevated_comment_data"`
																		Typename                                          string        `json:"__typename"`
																	} `json:"node"`
																	Cursor interface{} `json:"cursor"`
																} `json:"edges"`
																PageInfo struct {
																	EndCursor       string `json:"end_cursor"`
																	HasNextPage     bool   `json:"has_next_page"`
																	HasPreviousPage bool   `json:"has_previous_page"`
																	StartCursor     string `json:"start_cursor"`
																} `json:"page_info"`
															} `json:"comments"`
														} `json:"comment_rendering_instance_for_feed_location"`
														Id                            string      `json:"id"`
														ThreadingConfig               interface{} `json:"threading_config"`
														ShouldAlwaysShowComposerOnTop bool        `json:"should_always_show_composer_on_top"`
														TypingIndicatorRenderer       struct {
															Typename string `json:"__typename"`
															Feedback struct {
																Id string `json:"id"`
															} `json:"feedback"`
															ModuleOperationUseCometUFITopLevelCommentListComponentsFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_useCometUFITopLevelCommentListComponents_feedback"`
															ModuleComponentUseCometUFITopLevelCommentListComponentsFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_useCometUFITopLevelCommentListComponents_feedback"`
														} `json:"typing_indicator_renderer"`
														ViewerActor struct {
															Typename             string `json:"__typename"`
															Id                   string `json:"id"`
															IsActor              string `json:"__isActor"`
															Name                 string `json:"name"`
															ProfilePictureDepth0 struct {
																Uri string `json:"uri"`
															} `json:"profile_picture_depth_0"`
															ProfilePictureDepth1 struct {
																Uri string `json:"uri"`
															} `json:"profile_picture_depth_1"`
															Gender string `json:"gender"`
														} `json:"viewer_actor"`
														Url      string `json:"url"`
														Typename string `json:"__typename"`
														Plugins  []struct {
															Typename                                          string      `json:"__typename"`
															ContextId                                         interface{} `json:"context_id"`
															PostId                                            string      `json:"post_id,omitempty"`
															ModuleOperationUseCometUFIComposerPluginsFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_useCometUFIComposerPlugins_feedback"`
															ModuleComponentUseCometUFIComposerPluginsFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_useCometUFIComposerPlugins_feedback"`
															HasAvatar   bool   `json:"has_avatar,omitempty"`
															FeedbackId  string `json:"feedback_id,omitempty"`
															EmojiSize   int    `json:"emoji_size,omitempty"`
															ViewerActor struct {
																Typename string `json:"__typename"`
																Id       string `json:"id"`
															} `json:"viewer_actor,omitempty"`
															ShouldCondenseVideoPreview bool   `json:"should_condense_video_preview,omitempty"`
															OwningProfileId            string `json:"owning_profile_id,omitempty"`
														} `json:"plugins"`
														CommentComposerPlaceholder     string      `json:"comment_composer_placeholder"`
														CanViewerComment               bool        `json:"can_viewer_comment"`
														ConstituentBadgeBannerRenderer interface{} `json:"constituent_badge_banner_renderer"`
														HaveCommentsBeenDisabled       bool        `json:"have_comments_been_disabled"`
														AreLiveVideoCommentsDisabled   bool        `json:"are_live_video_comments_disabled"`
														IsViewerMuted                  bool        `json:"is_viewer_muted"`
														CommentsDisabledNoticeRenderer struct {
															Typename      string `json:"__typename"`
															NoticeMessage struct {
																DelightRanges     []interface{} `json:"delight_ranges"`
																ImageRanges       []interface{} `json:"image_ranges"`
																InlineStyleRanges []interface{} `json:"inline_style_ranges"`
																AggregatedRanges  []interface{} `json:"aggregated_ranges"`
																Ranges            []interface{} `json:"ranges"`
																ColorRanges       []interface{} `json:"color_ranges"`
																Text              string        `json:"text"`
															} `json:"notice_message"`
															ModuleOperationCometUFICommentDisabledNoticeFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_CometUFICommentDisabledNotice_feedback"`
															ModuleComponentCometUFICommentDisabledNoticeFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_CometUFICommentDisabledNotice_feedback"`
														} `json:"comments_disabled_notice_renderer"`
														WorkCommentSummariesFromFeedback interface{} `json:"work_comment_summaries_from_feedback"`
													} `json:"feedback"`
													ModuleOperationCometUFICommentListRendererFeedback struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_CometUFICommentListRenderer_feedback"`
													ModuleComponentCometUFICommentListRendererFeedback struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_CometUFICommentListRenderer_feedback"`
												} `json:"comment_list_renderer"`
												CometUfiSummaryAndActionsRenderer struct {
													Typename string `json:"__typename"`
													Feedback struct {
														Id                   string `json:"id"`
														SubscriptionTargetId string `json:"subscription_target_id"`
														I18NReactionCount    string `json:"i18n_reaction_count"`
														ImportantReactors    struct {
															Nodes []interface{} `json:"nodes"`
														} `json:"important_reactors"`
														ReactionCount struct {
															Count   int  `json:"count"`
															IsEmpty bool `json:"is_empty"`
														} `json:"reaction_count"`
														TopReactions struct {
															Count int `json:"count"`
															Edges []struct {
																VisibleInBlingBar bool `json:"visible_in_bling_bar"`
																Node              struct {
																	Id            string `json:"id"`
																	LocalizedName string `json:"localized_name"`
																} `json:"node"`
																I18NReactionCount string `json:"i18n_reaction_count"`
																ReactionCount     int    `json:"reaction_count"`
															} `json:"edges"`
														} `json:"top_reactions"`
														ReactionDisplayConfig struct {
															ReactionDisplayStrategy                       string      `json:"reaction_display_strategy"`
															ReactionStringWithViewer                      interface{} `json:"reaction_string_with_viewer"`
															ReactionStringWithoutViewer                   interface{} `json:"reaction_string_without_viewer"`
															ModuleOperationCometUFIReactionsCountFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_CometUFIReactionsCount_feedback"`
															ModuleComponentCometUFIReactionsCountFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_CometUFIReactionsCount_feedback"`
														} `json:"reaction_display_config"`
														ViewerActor struct {
															Typename string `json:"__typename"`
															Id       string `json:"id"`
															Name     string `json:"name"`
														} `json:"viewer_actor"`
														ViewerFeedbackReactionInfo        interface{} `json:"viewer_feedback_reaction_info"`
														CanShowSeenBy                     bool        `json:"can_show_seen_by"`
														IfViewerCanSeeSeenByMemberList    interface{} `json:"if_viewer_can_see_seen_by_member_list"`
														IfViewerCannotSeeSeenByMemberList struct {
															I18NReactionCount string `json:"i18n_reaction_count"`
															ReactionCount     struct {
																Count int `json:"count"`
															} `json:"reaction_count"`
															ReactionDisplayConfig struct {
																ReactionDisplayStrategy string `json:"reaction_display_strategy"`
															} `json:"reaction_display_config"`
															SeenBy struct {
																Count           interface{} `json:"count"`
																I18NSeenByCount interface{} `json:"i18n_seen_by_count"`
																SeenByEveryone  bool        `json:"seen_by_everyone"`
															} `json:"seen_by"`
															ModuleOperationCometUFISeenByCountFeedbackIfViewerCannotSeeSeenByMemberList struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_CometUFISeenByCount_feedback__if_viewer_cannot_see_seen_by_member_list"`
															ModuleComponentCometUFISeenByCountFeedbackIfViewerCannotSeeSeenByMemberList struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_CometUFISeenByCount_feedback__if_viewer_cannot_see_seen_by_member_list"`
															Id string `json:"id"`
														} `json:"if_viewer_cannot_see_seen_by_member_list"`
														I18NShareCount string `json:"i18n_share_count"`
														ShareCount     struct {
															Count   int  `json:"count"`
															IsEmpty bool `json:"is_empty"`
														} `json:"share_count"`
														CanSeeTopCustomReactionsForCreators interface{} `json:"can_see_top_custom_reactions_for_creators"`
														CommentsCountSummaryRenderer        struct {
															Typename string `json:"__typename"`
															Feedback struct {
																Id                string `json:"id"`
																TotalCommentCount int    `json:"total_comment_count"`
																I18NCommentCount  string `json:"i18n_comment_count"`
															} `json:"feedback"`
															ModuleOperationCometUFISummaryBaseFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_CometUFISummaryBase_feedback"`
															ModuleComponentCometUFISummaryBaseFeedback struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_CometUFISummaryBase_feedback"`
														} `json:"comments_count_summary_renderer"`
														AssociatedVideo struct {
															BroadcastIsAmaEnabled bool   `json:"broadcast_is_ama_enabled"`
															Id                    string `json:"id"`
														} `json:"associated_video"`
														TotalCommentCount      int         `json:"total_comment_count"`
														PagePrivateReply       interface{} `json:"page_private_reply"`
														VideoViewCount         int         `json:"video_view_count"`
														VideoViewCountRenderer interface{} `json:"video_view_count_renderer"`
														IsSimilarCqaQuestion   bool        `json:"is_similar_cqa_question"`
														MessageAction          interface{} `json:"message_action"`
														UfiActionRenderers     []struct {
															Typename string `json:"__typename"`
															Feedback struct {
																ViewerCustomReactionForCreator interface{} `json:"viewer_custom_reaction_for_creator"`
																ViewerFeedbackReactionInfo     interface{} `json:"viewer_feedback_reaction_info"`
																AssociatedVideo                struct {
																	Id string `json:"id"`
																} `json:"associated_video"`
																TopReactions struct {
																	Edges []struct {
																		ReactionCount int `json:"reaction_count"`
																		Node          struct {
																			Key int    `json:"key"`
																			Id  string `json:"id"`
																		} `json:"node"`
																	} `json:"edges"`
																} `json:"top_reactions"`
																Reactors struct {
																	Count   int  `json:"count"`
																	IsEmpty bool `json:"is_empty"`
																} `json:"reactors"`
																SupportedReactionInfos []struct {
																	Animation struct {
																		UriKeyframes2 string `json:"uri_keyframes2"`
																	} `json:"animation"`
																	Id string `json:"id"`
																} `json:"supported_reaction_infos"`
																IfViewerCanUseCreatorCustomReactions interface{} `json:"if_viewer_can_use_creator_custom_reactions"`
																Id                                   string      `json:"id"`
																ViewerActor                          struct {
																	Typename string `json:"__typename"`
																	Id       string `json:"id"`
																} `json:"viewer_actor"`
																IfViewerCanRenderCreatorCustomReactions interface{} `json:"if_viewer_can_render_creator_custom_reactions"`
															} `json:"feedback"`
															HideLabelForAMA                                                   bool `json:"hideLabelForAMA"`
															ModuleOperationUseCometUFIPostActionBarFeedbackUfiActionRenderers struct {
																Dr string `json:"__dr"`
															} `json:"__module_operation_useCometUFIPostActionBar_feedback__ufi_action_renderers"`
															ModuleComponentUseCometUFIPostActionBarFeedbackUfiActionRenderers struct {
																Dr string `json:"__dr"`
															} `json:"__module_component_useCometUFIPostActionBar_feedback__ufi_action_renderers"`
														} `json:"ufi_action_renderers"`
														ShouldShowReshareWarning bool `json:"should_show_reshare_warning"`
													} `json:"feedback"`
													ModuleOperationCometUFISummaryAndActionsFeedback struct {
														Dr string `json:"__dr"`
													} `json:"__module_operation_CometUFISummaryAndActions_feedback"`
													ModuleComponentCometUFISummaryAndActionsFeedback struct {
														Dr string `json:"__dr"`
													} `json:"__module_component_CometUFISummaryAndActions_feedback"`
												} `json:"comet_ufi_summary_and_actions_renderer"`
												OwningProfile struct {
													Typename string `json:"__typename"`
													Id       string `json:"id"`
												} `json:"owning_profile"`
												CommentModerationFilterRestrictionNotice interface{} `json:"comment_moderation_filter_restriction_notice"`
											} `json:"feedback"`
											ModuleOperationCometFeedUFIFeedback struct {
												Dr string `json:"__dr"`
											} `json:"__module_operation_CometFeedUFI_feedback"`
											ModuleComponentCometFeedUFIFeedback struct {
												Dr string `json:"__dr"`
											} `json:"__module_component_CometFeedUFI_feedback"`
										} `json:"ufi_renderer"`
										Id string `json:"id"`
									} `json:"feedback_target_with_context"`
									InterestingTopLevelComments []interface{} `json:"interesting_top_level_comments"`
								} `json:"feedback_context"`
								ShareableFromPerspectiveOfFeedUfi struct {
									Typename string `json:"__typename"`
									IsEntity string `json:"__isEntity"`
									Url      string `json:"url"`
									IsNode   string `json:"__isNode"`
									Id       string `json:"id"`
									PostId   string `json:"post_id"`
								} `json:"shareable_from_perspective_of_feed_ufi"`
								Id        string `json:"id"`
								Url       string `json:"url"`
								Shareable struct {
									Typename string `json:"__typename"`
									IsNode   string `json:"__isNode"`
									Id       string `json:"id"`
								} `json:"shareable"`
								SponsoredData               interface{}   `json:"sponsored_data"`
								InformTreatmentForMessaging interface{}   `json:"inform_treatment_for_messaging"`
								Tracking                    string        `json:"tracking"`
								PostId                      string        `json:"post_id"`
								VoteAttachments             []interface{} `json:"vote_attachments"`
							} `json:"story"`
							ModuleOperationCometFeedStoryFeedbackSectionStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_operation_CometFeedStoryFeedbackSection_story"`
							ModuleComponentCometFeedStoryFeedbackSectionStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_component_CometFeedStoryFeedbackSection_story"`
						} `json:"feedback"`
						OuterFooter  interface{} `json:"outer_footer"`
						CallToAction struct {
							Typename             string `json:"__typename"`
							IsICometStorySection string `json:"__isICometStorySection"`
							IsProdEligible       bool   `json:"is_prod_eligible"`
							Story                struct {
								Bumpers interface{} `json:"bumpers"`
								Id      string      `json:"id"`
							} `json:"story"`
							ModuleOperationCometFeedStoryCallToActionSectionStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_operation_CometFeedStoryCallToActionSection_story"`
							ModuleComponentCometFeedStoryCallToActionSectionStory struct {
								Dr string `json:"__dr"`
							} `json:"__module_component_CometFeedStoryCallToActionSection_story"`
						} `json:"call_to_action"`
						PostInformTreatment interface{} `json:"post_inform_treatment"`
					} `json:"comet_sections"`
					EncryptedTracking                                    string      `json:"encrypted_tracking"`
					ShouldHostActorLinkInWatch                           bool        `json:"should_host_actor_link_in_watch"`
					WhatsappAdContext                                    interface{} `json:"whatsapp_ad_context"`
					SchemaContext                                        interface{} `json:"schema_context"`
					ClickTrackingLinkshimCb                              string      `json:"click_tracking_linkshim_cb"`
					EncryptedClickTracking                               string      `json:"encrypted_click_tracking"`
					ModuleOperationCometFeedUnitContainerSectionFeedUnit struct {
						Dr string `json:"__dr"`
					} `json:"__module_operation_CometFeedUnitContainerSection_feedUnit"`
					ModuleComponentCometFeedUnitContainerSectionFeedUnit struct {
						Dr string `json:"__dr"`
					} `json:"__module_component_CometFeedUnitContainerSection_feedUnit"`
					IsTrackableFeedUnit string `json:"__isTrackableFeedUnit"`
					Trackingdata        struct {
						Id string `json:"id"`
					} `json:"trackingdata"`
					ViewabilityConfig []int `json:"viewability_config"`
					ClientViewConfig  struct {
						CanDelayLogImpression bool `json:"can_delay_log_impression"`
						UseBanzaiSignalImp    bool `json:"use_banzai_signal_imp"`
						UseBanzaiVitalImp     bool `json:"use_banzai_vital_imp"`
					} `json:"client_view_config"`
				} `json:"node"`
			} `json:"data"`
			Extensions struct {
				PrefetchUrisV2 []struct {
					Uri   string      `json:"uri"`
					Label interface{} `json:"label"`
				} `json:"prefetch_uris_v2"`
				AllVideoDashPrefetchRepresentations []struct {
					Representations []struct {
						RepresentationId      string `json:"representation_id"`
						MimeType              string `json:"mime_type"`
						Codecs                string `json:"codecs"`
						BaseUrl               string `json:"base_url"`
						Bandwidth             int    `json:"bandwidth"`
						Height                int    `json:"height"`
						Width                 int    `json:"width"`
						PlaybackResolutionMos string `json:"playback_resolution_mos"`
						Segments              []struct {
							Start int `json:"start"`
							End   int `json:"end"`
						} `json:"segments"`
					} `json:"representations"`
					VideoId int64 `json:"video_id"`
				} `json:"all_video_dash_prefetch_representations"`
				IsFinal bool `json:"is_final"`
			} `json:"extensions"`
		} `json:"result"`
		SequenceNumber int `json:"sequence_number"`
	} `json:"__bbox"`
}

type facebookReels struct {
	Bbox struct {
		Complete bool `json:"complete"`
		Result   struct {
			Data struct {
				Video struct {
					CreationStory struct {
						Tracking              string `json:"tracking"`
						ShortFormVideoContext struct {
							SelfViewBoost interface{} `json:"self_view_boost"`
							Video         struct {
								Id                string      `json:"id"`
								VideoCollaborator interface{} `json:"video_collaborator"`
								Embeddable        bool        `json:"embeddable"`
								Owner             struct {
									Typename string `json:"__typename"`
									Id       string `json:"id"`
								} `json:"owner"`
								AudioAvailability string `json:"audio_availability"`
							} `json:"video"`
							IfShouldChangeUrlForReels interface{} `json:"if_should_change_url_for_reels"`
							ShareableUrl              string      `json:"shareable_url"`
							PlaybackVideo             struct {
								Height                                 int           `json:"height"`
								Width                                  int           `json:"width"`
								LengthInSecond                         float64       `json:"length_in_second"`
								Id                                     string        `json:"id"`
								AnimatedImageCaption                   interface{}   `json:"animated_image_caption"`
								BroadcasterOrigin                      interface{}   `json:"broadcaster_origin"`
								BroadcastId                            interface{}   `json:"broadcast_id"`
								BroadcastStatus                        interface{}   `json:"broadcast_status"`
								IsLiveStreaming                        bool          `json:"is_live_streaming"`
								IsLiveTraceEnabled                     bool          `json:"is_live_trace_enabled"`
								IsLooping                              bool          `json:"is_looping"`
								IsVideoBroadcast                       bool          `json:"is_video_broadcast"`
								IsPodcastVideo                         bool          `json:"is_podcast_video"`
								LoopCount                              int           `json:"loop_count"`
								IsSpherical                            bool          `json:"is_spherical"`
								IsSphericalEnabled                     bool          `json:"is_spherical_enabled"`
								UnsupportedBrowserMessage              interface{}   `json:"unsupported_browser_message"`
								PmvMetadata                            interface{}   `json:"pmv_metadata"`
								LatencySensitiveConfig                 interface{}   `json:"latency_sensitive_config"`
								LivePlaybackInstrumentationConfigs     interface{}   `json:"live_playback_instrumentation_configs"`
								IsNcsr                                 bool          `json:"is_ncsr"`
								PermalinkUrl                           string        `json:"permalink_url"`
								CaptionsUrl                            interface{}   `json:"captions_url"`
								DashPrefetchExperimental               []string      `json:"dash_prefetch_experimental"`
								VideoAvailableCaptionsLocales          []interface{} `json:"video_available_captions_locales"`
								VideoStatusType                        string        `json:"video_status_type"`
								CanUseOz                               bool          `json:"can_use_oz"`
								DashManifest                           string        `json:"dash_manifest"`
								DashManifestUrl                        string        `json:"dash_manifest_url"`
								MinQualityPreference                   interface{}   `json:"min_quality_preference"`
								AudioUserPreferredLanguage             string        `json:"audio_user_preferred_language"`
								IsRssPodcastVideo                      bool          `json:"is_rss_podcast_video"`
								BrowserNativeSdUrl                     string        `json:"browser_native_sd_url"`
								BrowserNativeHdUrl                     string        `json:"browser_native_hd_url"`
								SphericalVideoFallbackUrls             interface{}   `json:"spherical_video_fallback_urls"`
								IsGamingVideo                          bool          `json:"is_gaming_video"`
								IsLatencyMenuEnabled                   bool          `json:"is_latency_menu_enabled"`
								FblsTier                               interface{}   `json:"fbls_tier"`
								IsLatencySensitiveBroadcast            bool          `json:"is_latency_sensitive_broadcast"`
								CometVideoPlayerStaticConfig           string        `json:"comet_video_player_static_config"`
								CometVideoPlayerContextSensitiveConfig string        `json:"comet_video_player_context_sensitive_config"`
								VideoPlayerShakaPerformanceLoggerInit  struct {
									Typename                                                          string `json:"__typename"`
									ModuleOperationUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
									ModuleComponentUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
								} `json:"video_player_shaka_performance_logger_init"`
								VideoPlayerShakaPerformanceLoggerShouldSample bool `json:"video_player_shaka_performance_logger_should_sample"`
								VideoPlayerShakaPerformanceLoggerInit2        struct {
									Typename                                                        string `json:"__typename"`
									ModuleOperationUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
									ModuleComponentUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
										Dr string `json:"__dr"`
									} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
									PerSessionSamplingRate interface{} `json:"per_session_sampling_rate"`
								} `json:"video_player_shaka_performance_logger_init2"`
								AutoplayGatingResult      string        `json:"autoplay_gating_result"`
								ViewerAutoplaySetting     string        `json:"viewer_autoplay_setting"`
								CanAutoplay               bool          `json:"can_autoplay"`
								DrmInfo                   string        `json:"drm_info"`
								P2PSettings               interface{}   `json:"p2p_settings"`
								AudioSettings             interface{}   `json:"audio_settings"`
								CaptionsSettings          interface{}   `json:"captions_settings"`
								BroadcastLowLatencyConfig interface{}   `json:"broadcast_low_latency_config"`
								AudioAvailability         string        `json:"audio_availability"`
								MutedSegments             []interface{} `json:"muted_segments"`
								SphericalVideoRenderer    interface{}   `json:"spherical_video_renderer"`
								PreferredThumbnail        struct {
									Image struct {
										Uri string `json:"uri"`
									} `json:"image"`
									ImagePreviewPayload string `json:"image_preview_payload"`
									Id                  string `json:"id"`
								} `json:"preferred_thumbnail"`
								VideoImfData          interface{} `json:"video_imf_data"`
								WarningScreenRenderer interface{} `json:"warning_screen_renderer"`
								CixScreen             interface{} `json:"cix_screen"`
							} `json:"playback_video"`
							VideoOwner struct {
								Typename               string `json:"__typename"`
								IsActor                string `json:"__isActor"`
								Id                     string `json:"id"`
								Name                   string `json:"name"`
								EnableReelsTabDeeplink bool   `json:"enable_reels_tab_deeplink"`
								IsVerified             bool   `json:"is_verified"`
								Url                    string `json:"url"`
								DisplayPicture         struct {
									Uri string `json:"uri"`
								} `json:"displayPicture"`
								SubscribeStatus string `json:"subscribe_status"`
							} `json:"video_owner"`
							IsPassiveContent       bool `json:"is_passive_content"`
							FbShortsReshareContext struct {
								IsReshare      bool `json:"is_reshare"`
								ReshareCreator struct {
									Typename               string `json:"__typename"`
									IsActor                string `json:"__isActor"`
									Id                     string `json:"id"`
									Name                   string `json:"name"`
									EnableReelsTabDeeplink bool   `json:"enable_reels_tab_deeplink"`
									IsVerified             bool   `json:"is_verified"`
									Url                    string `json:"url"`
								} `json:"reshare_creator"`
							} `json:"fb_shorts_reshare_context"`
							RemixInfo struct {
								IsRemixable bool   `json:"is_remixable"`
								Status      string `json:"status"`
							} `json:"remix_info"`
							VideoOwnerType string `json:"video_owner_type"`
							SoundtrackInfo struct {
								Id   string `json:"id"`
								Type string `json:"type"`
							} `json:"soundtrack_info"`
							TrackTitle                string `json:"track_title"`
							MusicAlbumArtUri          string `json:"music_album_art_uri"`
							IsOriginalAudioOnFacebook bool   `json:"is_original_audio_on_facebook"`
						} `json:"short_form_video_context"`
						Id                      string `json:"id"`
						EncryptedTracking       string `json:"encrypted_tracking"`
						EncryptedClickTracking  string `json:"encrypted_click_tracking"`
						ClickTrackingLinkshimCb string `json:"click_tracking_linkshim_cb"`
						PrivacyScope            struct {
							Description  string `json:"description"`
							Label        string `json:"label"`
							DisplayLabel string `json:"display_label"`
							IconImage    struct {
								Height int    `json:"height"`
								Scale  int    `json:"scale"`
								Uri    string `json:"uri"`
								Width  int    `json:"width"`
							} `json:"icon_image"`
							PrivacyScopeRenderer interface{} `json:"privacy_scope_renderer"`
						} `json:"privacy_scope"`
						CreationTime int `json:"creation_time"`
						Feedback     struct {
							AssociatedGroup interface{} `json:"associated_group"`
							Id              string      `json:"id"`
						} `json:"feedback"`
						BrandedContentPostInfo             interface{} `json:"branded_content_post_info"`
						PostId                             string      `json:"post_id"`
						CanViewerDelete                    bool        `json:"can_viewer_delete"`
						CanViewerCancelCollaborationInvite bool        `json:"can_viewer_cancel_collaboration_invite"`
						CanViewerRemoveCollaborator        bool        `json:"can_viewer_remove_collaborator"`
						CanViewerSeeCollaborationInvite    bool        `json:"can_viewer_see_collaboration_invite"`
						LegalReportingCtaType              interface{} `json:"legal_reporting_cta_type"`
						LegalReportingUri                  interface{} `json:"legal_reporting_uri"`
						SaveInfo                           struct {
							ViewerSaveState string `json:"viewer_save_state"`
						} `json:"save_info"`
						To    interface{} `json:"to"`
						Video struct {
							Id    string `json:"id"`
							Owner struct {
								Typename string `json:"__typename"`
								Id       string `json:"id"`
							} `json:"owner"`
						} `json:"video"`
						CollaboratorOrInvitees              []interface{} `json:"collaborator_or_invitees"`
						CanViewerRemoveSelfAsCollaborator   bool          `json:"can_viewer_remove_self_as_collaborator"`
						IfViewerCanSeeStarsToggleMenuOption interface{}   `json:"if_viewer_can_see_stars_toggle_menu_option"`
						SerializedFrtpIdentifiers           interface{}   `json:"serialized_frtp_identifiers"`
						DebugInfo                           interface{}   `json:"debug_info"`
						Message                             struct {
							Text   string `json:"text"`
							Ranges []struct {
								Offset int `json:"offset"`
								Length int `json:"length"`
								Entity struct {
									Typename  string `json:"__typename"`
									IsEntity  string `json:"__isEntity"`
									MobileUrl string `json:"mobileUrl"`
									Url       string `json:"url"`
									Id        string `json:"id"`
									IsNode    string `json:"__isNode"`
								} `json:"entity"`
							} `json:"ranges"`
						} `json:"message"`
						Actors []struct {
							Typename string `json:"__typename"`
							Name     string `json:"name"`
							Id       string `json:"id"`
						} `json:"actors"`
						BannerAdsOverlay                  interface{} `json:"banner_ads_overlay"`
						ContextualElementShortFormContext struct {
							VideoLabel interface{} `json:"video_label"`
						} `json:"contextualElementShortFormContext"`
					} `json:"creation_story"`
					Id string `json:"id"`
				} `json:"video"`
				Viewer struct {
					Actor struct {
						Typename string `json:"__typename"`
						Id       string `json:"id"`
					} `json:"actor"`
					LassoBlueFeed struct {
						Edges []struct {
							Node struct {
								Typename              string `json:"__typename"`
								Tracking              string `json:"tracking"`
								ShortFormVideoContext struct {
									SelfViewBoost interface{} `json:"self_view_boost"`
									Video         struct {
										Id                string      `json:"id"`
										VideoCollaborator interface{} `json:"video_collaborator"`
										Embeddable        bool        `json:"embeddable"`
										Owner             struct {
											Typename string `json:"__typename"`
											Id       string `json:"id"`
										} `json:"owner"`
										AudioAvailability string `json:"audio_availability"`
									} `json:"video"`
									IfShouldChangeUrlForReels struct {
										ShareableUrl string `json:"shareable_url"`
									} `json:"if_should_change_url_for_reels"`
									ShareableUrl  string `json:"shareable_url"`
									PlaybackVideo struct {
										Height                                 int           `json:"height"`
										Width                                  int           `json:"width"`
										LengthInSecond                         float64       `json:"length_in_second"`
										Id                                     string        `json:"id"`
										AnimatedImageCaption                   interface{}   `json:"animated_image_caption"`
										BroadcasterOrigin                      interface{}   `json:"broadcaster_origin"`
										BroadcastId                            interface{}   `json:"broadcast_id"`
										BroadcastStatus                        interface{}   `json:"broadcast_status"`
										IsLiveStreaming                        bool          `json:"is_live_streaming"`
										IsLiveTraceEnabled                     bool          `json:"is_live_trace_enabled"`
										IsLooping                              bool          `json:"is_looping"`
										IsVideoBroadcast                       bool          `json:"is_video_broadcast"`
										IsPodcastVideo                         bool          `json:"is_podcast_video"`
										LoopCount                              int           `json:"loop_count"`
										IsSpherical                            bool          `json:"is_spherical"`
										IsSphericalEnabled                     bool          `json:"is_spherical_enabled"`
										UnsupportedBrowserMessage              interface{}   `json:"unsupported_browser_message"`
										PmvMetadata                            interface{}   `json:"pmv_metadata"`
										LatencySensitiveConfig                 interface{}   `json:"latency_sensitive_config"`
										LivePlaybackInstrumentationConfigs     interface{}   `json:"live_playback_instrumentation_configs"`
										IsNcsr                                 bool          `json:"is_ncsr"`
										PermalinkUrl                           string        `json:"permalink_url"`
										CaptionsUrl                            interface{}   `json:"captions_url"`
										DashPrefetchExperimental               []string      `json:"dash_prefetch_experimental"`
										VideoAvailableCaptionsLocales          []interface{} `json:"video_available_captions_locales"`
										VideoStatusType                        string        `json:"video_status_type"`
										CanUseOz                               bool          `json:"can_use_oz"`
										DashManifest                           string        `json:"dash_manifest"`
										DashManifestUrl                        string        `json:"dash_manifest_url"`
										MinQualityPreference                   interface{}   `json:"min_quality_preference"`
										AudioUserPreferredLanguage             string        `json:"audio_user_preferred_language"`
										IsRssPodcastVideo                      bool          `json:"is_rss_podcast_video"`
										BrowserNativeSdUrl                     string        `json:"browser_native_sd_url"`
										BrowserNativeHdUrl                     string        `json:"browser_native_hd_url"`
										SphericalVideoFallbackUrls             interface{}   `json:"spherical_video_fallback_urls"`
										IsGamingVideo                          bool          `json:"is_gaming_video"`
										IsLatencyMenuEnabled                   bool          `json:"is_latency_menu_enabled"`
										FblsTier                               interface{}   `json:"fbls_tier"`
										IsLatencySensitiveBroadcast            bool          `json:"is_latency_sensitive_broadcast"`
										CometVideoPlayerStaticConfig           string        `json:"comet_video_player_static_config"`
										CometVideoPlayerContextSensitiveConfig string        `json:"comet_video_player_context_sensitive_config"`
										VideoPlayerShakaPerformanceLoggerInit  struct {
											Typename                                                          string `json:"__typename"`
											ModuleOperationUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
												Dr string `json:"__dr"`
											} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
											ModuleComponentUseVideoPlayerShakaPerformanceLoggerRelayImplVideo struct {
												Dr string `json:"__dr"`
											} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerRelayImpl_video"`
										} `json:"video_player_shaka_performance_logger_init"`
										VideoPlayerShakaPerformanceLoggerShouldSample bool `json:"video_player_shaka_performance_logger_should_sample"`
										VideoPlayerShakaPerformanceLoggerInit2        struct {
											Typename                                                        string `json:"__typename"`
											ModuleOperationUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
												Dr string `json:"__dr"`
											} `json:"__module_operation_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
											ModuleComponentUseVideoPlayerShakaPerformanceLoggerBuilderVideo struct {
												Dr string `json:"__dr"`
											} `json:"__module_component_useVideoPlayerShakaPerformanceLoggerBuilder_video"`
											PerSessionSamplingRate interface{} `json:"per_session_sampling_rate"`
										} `json:"video_player_shaka_performance_logger_init2"`
										AutoplayGatingResult      string        `json:"autoplay_gating_result"`
										ViewerAutoplaySetting     string        `json:"viewer_autoplay_setting"`
										CanAutoplay               bool          `json:"can_autoplay"`
										DrmInfo                   string        `json:"drm_info"`
										P2PSettings               interface{}   `json:"p2p_settings"`
										AudioSettings             interface{}   `json:"audio_settings"`
										CaptionsSettings          interface{}   `json:"captions_settings"`
										BroadcastLowLatencyConfig interface{}   `json:"broadcast_low_latency_config"`
										AudioAvailability         string        `json:"audio_availability"`
										MutedSegments             []interface{} `json:"muted_segments"`
										SphericalVideoRenderer    interface{}   `json:"spherical_video_renderer"`
										PreferredThumbnail        struct {
											Image struct {
												Uri string `json:"uri"`
											} `json:"image"`
											ImagePreviewPayload interface{} `json:"image_preview_payload"`
											Id                  string      `json:"id"`
										} `json:"preferred_thumbnail"`
										VideoImfData          interface{} `json:"video_imf_data"`
										WarningScreenRenderer interface{} `json:"warning_screen_renderer"`
										CixScreen             interface{} `json:"cix_screen"`
									} `json:"playback_video"`
									VideoOwner struct {
										Typename       string `json:"__typename"`
										IsActor        string `json:"__isActor"`
										Id             string `json:"id"`
										Name           string `json:"name"`
										Username       string `json:"username"`
										Url            string `json:"url"`
										DisplayPicture struct {
											Uri string `json:"uri"`
										} `json:"displayPicture"`
									} `json:"video_owner"`
									IsPassiveContent       bool        `json:"is_passive_content"`
									FbShortsReshareContext interface{} `json:"fb_shorts_reshare_context"`
									RemixInfo              struct {
										IsRemixable bool   `json:"is_remixable"`
										Status      string `json:"status"`
									} `json:"remix_info"`
									VideoOwnerType string `json:"video_owner_type"`
									SoundtrackInfo struct {
										Id   string `json:"id"`
										Type string `json:"type"`
									} `json:"soundtrack_info"`
									TrackTitle                string `json:"track_title"`
									MusicAlbumArtUri          string `json:"music_album_art_uri"`
									IsOriginalAudioOnFacebook bool   `json:"is_original_audio_on_facebook"`
								} `json:"short_form_video_context"`
								Id                      string      `json:"id"`
								EncryptedTracking       string      `json:"encrypted_tracking"`
								EncryptedClickTracking  string      `json:"encrypted_click_tracking"`
								ClickTrackingLinkshimCb string      `json:"click_tracking_linkshim_cb"`
								PrivacyScope            interface{} `json:"privacy_scope"`
								CreationTime            int         `json:"creation_time"`
								Feedback                struct {
									AssociatedGroup interface{} `json:"associated_group"`
									Id              string      `json:"id"`
								} `json:"feedback"`
								BrandedContentPostInfo             interface{} `json:"branded_content_post_info"`
								PostId                             string      `json:"post_id"`
								CanViewerDelete                    bool        `json:"can_viewer_delete"`
								CanViewerCancelCollaborationInvite bool        `json:"can_viewer_cancel_collaboration_invite"`
								CanViewerRemoveCollaborator        bool        `json:"can_viewer_remove_collaborator"`
								CanViewerSeeCollaborationInvite    bool        `json:"can_viewer_see_collaboration_invite"`
								LegalReportingCtaType              interface{} `json:"legal_reporting_cta_type"`
								LegalReportingUri                  interface{} `json:"legal_reporting_uri"`
								SaveInfo                           struct {
									ViewerSaveState string `json:"viewer_save_state"`
								} `json:"save_info"`
								To    interface{} `json:"to"`
								Video struct {
									Id    string `json:"id"`
									Owner struct {
										Typename string `json:"__typename"`
										Id       string `json:"id"`
									} `json:"owner"`
								} `json:"video"`
								CollaboratorOrInvitees              []interface{} `json:"collaborator_or_invitees"`
								CanViewerRemoveSelfAsCollaborator   bool          `json:"can_viewer_remove_self_as_collaborator"`
								IfViewerCanSeeStarsToggleMenuOption interface{}   `json:"if_viewer_can_see_stars_toggle_menu_option"`
								SerializedFrtpIdentifiers           interface{}   `json:"serialized_frtp_identifiers"`
								DebugInfo                           interface{}   `json:"debug_info"`
								Message                             struct {
									Text   string `json:"text"`
									Ranges []struct {
										Offset int `json:"offset"`
										Length int `json:"length"`
										Entity struct {
											Typename  string `json:"__typename"`
											IsEntity  string `json:"__isEntity"`
											MobileUrl string `json:"mobileUrl"`
											Url       string `json:"url"`
											IsNode    string `json:"__isNode"`
											Id        string `json:"id"`
										} `json:"entity"`
									} `json:"ranges"`
								} `json:"message"`
								Actors []struct {
									Typename string `json:"__typename"`
									Name     string `json:"name"`
									Id       string `json:"id"`
								} `json:"actors"`
								BannerAdsOverlay                  interface{} `json:"banner_ads_overlay"`
								ContextualElementShortFormContext struct {
									VideoLabel interface{} `json:"video_label"`
								} `json:"contextualElementShortFormContext"`
								IsNode string `json:"__isNode"`
							} `json:"node"`
							Cursor string `json:"cursor"`
							Id     string `json:"id"`
						} `json:"edges"`
						PageInfo struct {
							EndCursor   string `json:"end_cursor"`
							HasNextPage bool   `json:"has_next_page"`
						} `json:"page_info"`
					} `json:"lasso_blue_feed"`
				} `json:"viewer"`
				QpViewer struct {
					EligiblePromotions struct {
						Nodes []interface{} `json:"nodes"`
					} `json:"eligible_promotions"`
				} `json:"qp_viewer"`
			} `json:"data"`
			Extensions struct {
				PrefetchUrisV2 []struct {
					Uri   string      `json:"uri"`
					Label interface{} `json:"label"`
				} `json:"prefetch_uris_v2"`
				AllVideoDashPrefetchRepresentations []struct {
					Representations []struct {
						RepresentationId      string `json:"representation_id"`
						MimeType              string `json:"mime_type"`
						Codecs                string `json:"codecs"`
						BaseUrl               string `json:"base_url"`
						Bandwidth             int    `json:"bandwidth"`
						Height                int    `json:"height"`
						Width                 int    `json:"width"`
						PlaybackResolutionMos string `json:"playback_resolution_mos"`
						Segments              []struct {
							Start int `json:"start"`
							End   int `json:"end"`
						} `json:"segments"`
					} `json:"representations"`
					VideoId int64 `json:"video_id"`
				} `json:"all_video_dash_prefetch_representations"`
				IsFinal bool `json:"is_final"`
			} `json:"extensions"`
		} `json:"result"`
		SequenceNumber int `json:"sequence_number"`
	} `json:"__bbox"`
}

type videoTitle struct {
	Text string `json:"text"`
}

func getFinallyUrl(link string) (error, string) {
	lastURL := link
	// 创建一个自定义的 Client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 获取最终跳转的地址
			lastURL = req.URL.String()
			return nil
		},
	}

	// 创建一个新的 GET 请求
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err, lastURL
	}

	// 设置请求头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Cookie", "sb=GLkAZJQdYfckMxH6lECElWSg; c_user=100026471882579; m_ls=%7B%22c%22%3A%7B%221%22%3A%22HCwAABZEFqqu8bYBEwUWpv2Ypae-LQA%22%2C%222%22%3A%22GSwVQBxMAAAWABbA0I7MDBYAABV-HEwAABYAFsbQjswMFgAAFigA%22%2C%2295%22%3A%22HCwAABYEForuqJsDEwUWpv2Ypae-LQA%22%7D%2C%22d%22%3A%22daea6326-862c-4b75-a50f-7933bb4cf1fb%22%2C%22s%22%3A%220%22%2C%22u%22%3A%22s0e9j3%22%7D; fbl_cs=AhDlAjyEZhqQ9pD3Sd1LpC%2FGGD0xOUI3cTZ0dmlRWDkrPXlzaWhpcmkxcQ; fbl_ci=687628016827021; vpd=v1%3B844x390x3; fbl_st=100639923%3BT%3A28417173; wl_cbv=v2%3Bclient_version%3A2392%3Btimestamp%3A1705030393; ps_n=0; wd=1728x869; datr=VenfZcjQMyGd_sZ2q_gyUOD9; xs=23%3Az-GSUb6sfSpEPA%3A2%3A1677769042%3A-1%3A-1%3A%3AAcV6kzthjN6camXHusY0d_b46CyLBvW2snq6QLXTTsM; fr=1J6AWE6wV1IH1inoo.AWUaKpkMzO69JHeczLefZulA1iU.Bl3-lX.ms.AAA.0.0.Bl3-lX.AWWqP3Wnb5I; presence=C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A1709173088316%2C%22v%22%3A1%7D")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err, lastURL
	}
	defer resp.Body.Close()

	return nil, lastURL
}

// https://www.facebook.com/reel/261735130226721?mibextid=9drbnH
// https://www.facebook.com/share/r/v5VgMiY5H3VUsPMp/?mibextid=KsPBc6

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {

	if !strings.Contains(url, `/reel/`) && !strings.Contains(url, `/watch/`) {
		_, url = getFinallyUrl(url)
	}
	var isPoster bool
	posterIds := utils.MatchOneOf(url, `(\d+)/posts/(\d+)`)
	if posterIds != nil && len(posterIds) >= 3 {
		isPoster = true
		url = "https://www.facebook.com/story.php/?id=" + posterIds[1] + "&story_fbid=" + posterIds[2]
	} else {
		vid := utils.MatchOneOf(url, `/reel/(\d+)`)
		if vid != nil && len(vid) >= 2 {
			url = "https://www.facebook.com/watch/?v=" + vid[1]
		}
	}

	var err error
	html, err := request.Get(url, url, map[string]string{
		"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Cookie":         "sb=GLkAZJQdYfckMxH6lECElWSg; c_user=100026471882579; m_ls=%7B%22c%22%3A%7B%221%22%3A%22HCwAABZEFqqu8bYBEwUWpv2Ypae-LQA%22%2C%222%22%3A%22GSwVQBxMAAAWABbA0I7MDBYAABV-HEwAABYAFsbQjswMFgAAFigA%22%2C%2295%22%3A%22HCwAABYEForuqJsDEwUWpv2Ypae-LQA%22%7D%2C%22d%22%3A%22daea6326-862c-4b75-a50f-7933bb4cf1fb%22%2C%22s%22%3A%220%22%2C%22u%22%3A%22s0e9j3%22%7D; fbl_cs=AhDlAjyEZhqQ9pD3Sd1LpC%2FGGD0xOUI3cTZ0dmlRWDkrPXlzaWhpcmkxcQ; fbl_ci=687628016827021; vpd=v1%3B844x390x3; fbl_st=100639923%3BT%3A28417173; wl_cbv=v2%3Bclient_version%3A2392%3Btimestamp%3A1705030393; ps_n=0; wd=1728x869; datr=VenfZcjQMyGd_sZ2q_gyUOD9; xs=23%3Az-GSUb6sfSpEPA%3A2%3A1677769042%3A-1%3A-1%3A%3AAcV6kzthjN6camXHusY0d_b46CyLBvW2snq6QLXTTsM; fr=1J6AWE6wV1IH1inoo.AWUaKpkMzO69JHeczLefZulA1iU.Bl3-lX.ms.AAA.0.0.Bl3-lX.AWWqP3Wnb5I; presence=C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A1709173088316%2C%22v%22%3A1%7D",
		"Sec-Fetch-Site": "none",
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dataStrings := utils.MatchOneOf(html, `[adp_CometVideoHomeNewPermalinkHeroUnitQueryRelayPreloader|adp_FBReelsRootWithEntrypointQueryRelayPreloader]\w+",(\{"__bbox":\{.*"browser_native_sd_url".*\}\})\]`)
	if dataStrings == nil || len(dataStrings) < 2 {
		fmt.Println(html)
		return nil, errors2.Annotate(errors.New("未匹配到视频数据"), "未匹配到视频数据")
	}

	if isPoster {
		var data facebookPoster
		if err = json.Unmarshal([]byte(dataStrings[1]), &data); err != nil {
			return nil, errors.WithStack(err)
		}

		var isNotVideo bool
		var images []string

		streams := make(map[string]*extractors.Stream, 2)

		streams["sd"] = &extractors.Stream{
			Parts: []*extractors.Part{
				{
					URL:  data.Bbox.Result.Data.Node.CometSections.Content.Story.Attachments[0].Styles.Attachment.Media.BrowserNativeSdUrl,
					Size: 850,
				},
			},
			Size:    360,
			Quality: "sd",
		}

		if len(data.Bbox.Result.Data.Node.CometSections.Content.Story.Attachments[0].Styles.Attachment.Media.BrowserNativeHdUrl) > 0 {
			streams["hd"] = &extractors.Stream{
				Parts: []*extractors.Part{
					{
						URL:  data.Bbox.Result.Data.Node.CometSections.Content.Story.Attachments[0].Styles.Attachment.Media.BrowserNativeHdUrl,
						Size: 850,
					},
				},
				Size:    1080,
				Quality: "hd",
			}
		}

		return []*extractors.Data{
			{
				Site:       "Facebook facebook.com",
				Title:      data.Bbox.Result.Data.Node.CometSections.Content.Story.Message.Text,
				Streams:    streams,
				URL:        url,
				Image:      images,
				Cover:      data.Bbox.Result.Data.Node.CometSections.Content.Story.Attachments[0].Styles.Attachment.Media.ThumbnailImage.Uri,
				IsNotVideo: isNotVideo,
			},
		}, nil
	} else {

		var data facebook
		if err = json.Unmarshal([]byte(dataStrings[1]), &data); err != nil {
			return nil, errors.WithStack(err)
		}

		if data.Bbox.Result.Data.Video.Story.Attachments == nil {
			var dataReel facebookReels
			if err = json.Unmarshal([]byte(dataStrings[1]), &dataReel); err != nil {
				return nil, errors.WithStack(err)
			}

			var isNotVideo bool
			var images []string

			streams := make(map[string]*extractors.Stream, 2)

			streams["sd"] = &extractors.Stream{
				Parts: []*extractors.Part{
					{
						URL:  dataReel.Bbox.Result.Data.Video.CreationStory.ShortFormVideoContext.PlaybackVideo.BrowserNativeSdUrl,
						Size: 850,
					},
				},
				Size:    360,
				Quality: "sd",
			}

			if len(dataReel.Bbox.Result.Data.Video.CreationStory.ShortFormVideoContext.PlaybackVideo.BrowserNativeHdUrl) > 0 {
				streams["hd"] = &extractors.Stream{
					Parts: []*extractors.Part{
						{
							URL:  dataReel.Bbox.Result.Data.Video.CreationStory.ShortFormVideoContext.PlaybackVideo.BrowserNativeHdUrl,
							Size: 850,
						},
					},
					Size:    1080,
					Quality: "hd",
				}
			}

			return []*extractors.Data{
				{
					Site:       "Facebook facebook.com",
					Title:      dataReel.Bbox.Result.Data.Video.CreationStory.Message.Text,
					Streams:    streams,
					URL:        url,
					Image:      images,
					Cover:      dataReel.Bbox.Result.Data.Video.CreationStory.ShortFormVideoContext.PlaybackVideo.PreferredThumbnail.Image.Uri,
					IsNotVideo: isNotVideo,
				},
			}, nil
		} else {
			titleStrings := utils.MatchOneOf(html, `"story":.*\{"message".*","delight_ranges`)

			if titleStrings == nil || len(titleStrings) < 1 {
				return nil, errors2.Annotate(errors.New("未匹配到视频标题"), "未匹配到视频标题")
			}

			titleStrings = utils.MatchOneOf(titleStrings[0], `"text".*","`)
			if titleStrings == nil || len(titleStrings) < 1 {
				return nil, errors2.Annotate(errors.New("未匹配到视频标题"), "未匹配到视频标题")
			}

			titleString := strings.Replace(titleStrings[0], `,"`, "", -1)
			var title videoTitle
			if err = json.Unmarshal([]byte("{"+titleString+"}"), &title); err != nil {
				return nil, errors.WithStack(err)
			}

			var isNotVideo bool
			var images []string

			streams := make(map[string]*extractors.Stream, 2)

			streams["sd"] = &extractors.Stream{
				Parts: []*extractors.Part{
					{
						URL:  data.Bbox.Result.Data.Video.Story.Attachments[0].Media.BrowserNativeSdUrl,
						Size: 850,
					},
				},
				Size:    360,
				Quality: "sd",
			}

			if len(data.Bbox.Result.Data.Video.Story.Attachments[0].Media.BrowserNativeHdUrl) > 0 {
				streams["hd"] = &extractors.Stream{
					Parts: []*extractors.Part{
						{
							URL:  data.Bbox.Result.Data.Video.Story.Attachments[0].Media.BrowserNativeHdUrl,
							Size: 850,
						},
					},
					Size:    1080,
					Quality: "hd",
				}
			}

			return []*extractors.Data{
				{
					Site:       "Facebook facebook.com",
					Title:      strings.Replace(title.Text, `\n`, "", -1),
					Streams:    streams,
					URL:        url,
					Image:      images,
					Cover:      data.Bbox.Result.Data.Video.Story.Attachments[0].Media.PreferredThumbnail.Image.Uri,
					IsNotVideo: isNotVideo,
				},
			}, nil
		}
	}
}
