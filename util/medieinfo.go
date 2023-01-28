package util

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/log"
	"io"
	"os/exec"
)

type MediaInfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                  string `json:"@type"`
			VideoCount            string `json:"VideoCount,omitempty"`
			FileExtension         string `json:"FileExtension,omitempty"`
			Format                string `json:"Format"`
			FormatProfile         string `json:"Format_Profile"`
			CodecID               string `json:"CodecID"`
			CodecIDCompatible     string `json:"CodecID_Compatible,omitempty"`
			FileSize              string `json:"FileSize,omitempty"`
			Duration              string `json:"Duration"`
			OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRate        string `json:"OverallBitRate,omitempty"`
			FrameRate             string `json:"FrameRate"`
			FrameCount            string `json:"FrameCount"`
			StreamSize            string `json:"StreamSize"`
			HeaderSize            string `json:"HeaderSize,omitempty"`
			DataSize              string `json:"DataSize,omitempty"`
			FooterSize            string `json:"FooterSize,omitempty"`
			IsStreamable          string `json:"IsStreamable,omitempty"`
			EncodedDate           string `json:"Encoded_Date"`
			TaggedDate            string `json:"Tagged_Date"`
			FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
			Extra                 struct {
				TIM                   string `json:"TIM,omitempty"`
				TSC                   string `json:"TSC,omitempty"`
				TSZ                   string `json:"TSZ,omitempty"`
				CodecConfigurationBox string `json:"CodecConfigurationBox,omitempty"`
			} `json:"extra"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
			BitRateMode                    string `json:"BitRate_Mode,omitempty"`
			BitRate                        string `json:"BitRate,omitempty"`
			Width                          string `json:"Width,omitempty"`
			Height                         string `json:"Height,omitempty"`
			StoredWidth                    string `json:"Stored_Width,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			Rotation                       string `json:"Rotation,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			ScanType                       string `json:"ScanType,omitempty"`
			Language                       string `json:"Language,omitempty"`
			BufferSize                     string `json:"BufferSize,omitempty"`
			ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
			ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
			ColourRange                    string `json:"colour_range,omitempty"`
			ColourRangeSource              string `json:"colour_range_Source,omitempty"`
			ColourPrimaries                string `json:"colour_primaries,omitempty"`
			ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
			TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
			TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
			MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
			MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

func GetMediaInfo(absPath string) (MediaInfo, error) {
	var md MediaInfo
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Panicf("cmd.Run产生的错误:%v\n", err)
	}

	//读取所有输出
	bytes, err := io.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return MediaInfo{}, err
	} else {
		if err := json.Unmarshal(bytes, &md); err != nil {
			return MediaInfo{}, err
		}
	}

	if err = cmd.Wait(); err != nil {
		log.Debug.Panicf("命令执行中有错误产生:%v\n", err)
	}
	return md, nil
}
