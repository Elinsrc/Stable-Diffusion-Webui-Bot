package i18n

func init() {
	register(text{language: ko_kr, Code: "ko_kr", Name: "한국인"})
}

var ko_kr = map[string]string{
	"help":                  "메시지 상자에 키워드 태그를 보내기만 하면 됩니다(영어로만 제공)\n태그 형식:loli,white tights,Uniform\n[동시에] 사진을 보낼 때 태그를 추가 할 수도 있습니다 (태그는 사진을 보낼 때 설명에 작성)\n\n{}를 사용하지 않고 ()를 사용, ()는 가중치를 1.1배 증가시키고 []는 가중치를 1.1배 감소시킵니다.\n사용법: (masterpiece:1.1), ((best quality)), some tags...\n\n사진 옵션:\nFT: 미세 조정(숫자가 클수록 변경 정도가 커짐)\nSPR: 초고해상도(숫자는 사진의 배율을 나타냄)",
	"history":               "역사적인 사진을 보려면 로그인하려면 웹 사이트로 이동하십시오.",
	"setLangSuccess":        "언어를 한국어로 설정",
	"cancel":                "취소",
	"size":                  "사진의 크기",
	"number":                "사진 수",
	"mode":                  "샘플링 모드",
	"unwanted":              "원하지 않는",
	"confirm":               "확인하다",
	"taskExist":             "기존 작업이 있습니다. 현재 작업이 끝날 때까지 기다리십시오",
	"generating":            "생성 중(메시지가 사라지지 않으면 생성 중임을 의미)\n잠시 기다려 주십시오. 생성에 실패하면 다시 시도하십시오",
	"joinGroup":             "현재 사용량이 제한되어 있습니다. 그룹에 가입하여 제한을 해제하세요.",
	"customUC":              "표시를 원하지 않는 내용을 보내주세요(역태그):",
	"nsfw":                  "어린이에게 적합하지 않음(Nsfw)",
	"lowQuality":            "저품질",
	"badAnatomy":            "잘못된 구조",
	"none":                  "없음(선택 취소)",
	"custom":                "사용자 정의",
	"strength":              "힘",
	"strengthInfo":          "업로드된 이미지가 변경되는 정도를 제어합니다. 강도를 낮추면 원본에 더 가까운 이미지를 생성합니다",
	"serErr":                "서버에 오류가 발생했습니다. 나중에 시도하십시오. 또는 토론과 도움을 위해 그룹에 가입하십시오.",
	"prohibit":              "오늘의 무료 사용 한도에 도달했습니다. 한 달 동안 무제한으로 봇을 계속 사용하려면 3$ 이상의 금액으로 후원해 주세요.\n일일 한도는 {{.time}} 후에 재설정됩니다.",
	"freeTimes":             "오늘 남은 시간 (그룹에 가입하면 무료 사용 횟수가 늘어날 수 있습니다): ",
	"clickMe":               "점프하려면 클릭",
	"translation":           "번역하다",
	"translate":             "영어로 자동 번역되는 태그",
	"reDraw":                "재생하다",
	"sendTag":               "태그를 보내주세요:",
	"sendPhoto":             "원본 사진을 보내주세요(압축하지 마세요):",
	"parsePhotoErr":         "파싱 ​​실패, 가장 원본 사진을 보내주세요(압축하지 마세요)...",
	"privateChat":           "로봇과 개인적으로 채팅하십시오",
	"tokenErr":              "Token 유효하지 않은",
	"model":                 "모델",
	"scale":                 "적당",
	"scaleInfo":             "이미지가 태그를 얼마나 강력하게 준수해야 하는지 - 값이 낮을수록 더 창의적인 결과를 생성합니다",
	"steps":                 "걸음 수",
	"stepsInfo":             "반복 횟수 - 값이 높을수록 빌드 시간이 길어지고 잠재적으로 더 상세하고 깔끔한 결과(및 잠재적으로 더 나쁜 결과)가 발생합니다",
	"sendImg":               "사진을 보내주세요(총 해상도 W*H는 4194304를 초과할 수 없음):",
	"bigImg":                "이미지 해상도가 너무 큽니다(총 해상도 W*H는 4194304를 초과할 수 없음).",
	"magnification":         "배율 선택",
	"edit":                  "편집하다",
	"modelInfo":             "모델마다 큰 차이가 있습니다. 그림 스타일, 캐릭터, 풍경, 크기 등에 많은 차이가 있습니다.",
	"modeInfo":              "모드에 따라 기본 페인팅 스타일과 내용에 영향을 주지 않고 속도와 결과에 약간의 차이가 있습니다.",
	"ucInfo":                "원하지 않는 콘텐츠, 일반적으로 태그의 반의어",
	"wait":                  "전송 후 잠시 기다립니다.",
	"dontDelMsg":            "이 메시지를 삭제하지 마십시오",
	"editTag":               "태그 편집",
	"Happend":               "머리 삽입",
	"Eappend":               "꼬리 삽입",
	"setImg":                "이미지 설정",
	"setImgInfo":            "업로드된 이미지를 바탕으로 그리기",
	"clearImg":              "선명한 이미지",
	"mustShare":             "구독한 사용자만 이 설정을 수정할 권한이 있습니다. 구독을 구매하려면 이동하세요.",
	"enable":                "활성화",
	"disable":               "장애가 있는",
	"reset":                 "초기화",
	"shareInfo":             "이 옵션은 결과 이미지를 웹 사이트에서 공유할지 여부를 결정하며 구독 취소자는 항상 이미지를 공유합니다.",
	"resetSeed":             "시드 재설정",
	"extraModel":            "추가 모델",
	"switch":                "스위치",
	"extraModelInfo":        "추가 모델은 관련 모델을 활성화하기 위한 것일 뿐, 관련 태그가 추가되지 않으면 무효가 됩니다. 예를 들어 Nahida 모델이 로드된 경우 적용하려면 태그에 nahida를 추가해야 합니다.",
	"noSubscribe":           "현재 활성 구독이 없거나 구독이 만료되었습니다.",
	"setControl":            "제어 그림",
	"editControl":           "관리도 편집",
	"delControl":            "컨트롤 차트 삭제",
	"controlPreprocess":     "전처리기",
	"controlProcess":        "프로세서",
	"back":                  "<<< 뒤쪽에",
	"setDft":                "기본값으로 설정",
	"onlySubscribe":         "구독한 사용자만 사용할 수 있습니다. 구독을 구매하려면 이동하세요.",
	"canny":                 "에지 감지",
	"depth":                 "깊이 추정-MiDaS",
	"depth_leres":           "깊이 추정-LeReS",
	"hed":                   "소프트 에지 검출-HED",
	"hed_safe":              "안전한 소프트 에지 검출-HED",
	"mediapipe_face":        "얼굴 가장자리 검출",
	"mlsd":                  "직선 선분 검출-M-LSD",
	"normal_map":            "법선 맵 추출-Midas",
	"openpose":              "자세 추정-OpenPose",
	"openpose_hand":         "자세 추정 | 손-OpenPose",
	"openpose_face":         "자세 추정 | 얼굴-OpenPose",
	"openpose_faceonly":     "얼굴만-OpenPose",
	"openpose_full":         "자세 추정 | 손 | 얼굴-OpenPose",
	"clip_vision":           "스타일 변환 처리-자동 적응",
	"color":                 "색상 픽셀화 처리-자동 적응",
	"pidinet":               "안전한 소프트 에지 검출-PiDiNet",
	"pidinet_safe":          "안전한 소프트 에지 검출-PiDiNet",
	"pidinet_sketch":        "손그림 에지 처리-자동 적응",
	"pidinet_scribble":      "낙서-손그림",
	"scribble_xdog":         "낙서-강조된 에지",
	"scribble_hed":          "낙서-합성",
	"threshold":             "임계 값",
	"depth_zoe":             "깊이 추정-ZoE",
	"normal_bae":            "법선 맵 추출-Bae",
	"oneformer_coco":        "시맨틱 분할-OneFormer-COCO",
	"oneformer_ade20k":      "시맨틱 분할-OneFormer-ADE20K",
	"lineart":               "선화 추출",
	"lineart_coarse":        "거친 선화 추출",
	"lineart_anime":         "애니메이션 선화 추출",
	"lineart_standard":      "표준 선화 추출-반전",
	"shuffle":               "랜덤 셔플",
	"tile_gaussian":         "블록 재샘플링",
	"invert":                "색상 반전",
	"lineart_anime_denoise": "애니메이션 선화 추출-노이즈 제거",
	"reference_only":        "참조 입력만",
	"inpaint":               "재그림-글로벌 퓨전 알고리즘",
	"invite":                "새로운 사용자(사용한 적이 없거나 사용 시간이 15분을 초과하지 않는 사용자)가 귀하의 초대 링크를 클릭한 후\n5회 추가 사용 기회를 얻을 수 있습니다(재설정되지 않으며 누적 가능)\n초대받은 새로운 사용자는 10회 추가 사용 횟수를 얻을 수 있습니다!",
	"inviteSuccess":         "사용자 {{user}}를 성공적으로 초대하셨습니다\n5회 추가 사용 횟수를 얻었으며\n남은 사용 횟수: {{freeAmount}}",
	"wasInvited":            "사용자 {{user}}에 의해 초대되었습니다. 10회 추가 사용 횟수를 얻었으며\n남은 사용 횟수: {{freeAmount}}",
	"freeMaxNum":            "무료 사용자는 매번 최대 3장의 사진을 생성할 수 있습니다.",
}
