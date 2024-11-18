package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"_UP_A4A_11_=wb962111be974404b955b13f14ada127; b-user-id=0a190155-ef43-1be6-890d-4681afd6837b; _UP_30C_6A_=st96b620196r4s1h44x0cg0zv91i3acv; _UP_TS_=sg1922d58b0e3edc0e8b2f9adc2ef4c91c2; _UP_E37_B7_=sg1922d58b0e3edc0e8b2f9adc2ef4c91c2; _UP_TG_=st96b620196r4s1h44x0cg0zv91i3acv; _UP_335_2B_=1; isg=BC0t70vUNqhvodJE9mBdhARbMsmnimFcm4wkbG8xXUQ45kGYN9oaLFn_1LoA_XkU; tfstk=ff8qFTVOyq34MG0_5s7NT3TAa49AxZDCnF61sCAGGtXc5mKw4COEM19XsO-NECB0StDOQav1tlxbmsVlUTvkGr61oOVwCCtXQChAsOvMQOtjV2OvMNQidFuIRIE-wfXWbibgEffGtt_HQWdvMNF8NRmBqIeNp6XJnNvir_fhEs4GIdblqs6uo14cs8cPe1bGsNj0E_fNg1fgodAoa1BlI3RvSeiA6icI4KtFr2EEW9RcUyRJzsDPC2BaSnTV_GrvgTSNbU5V09sRu-JDoLOMlEI8abb9NHJwjpD7ctxF_BtwL24Ownv2rG8S62dAutJGb3HURs7X4M82Dfekip62xZd_TxOPj3TAvEMuuwxpGZsPORCoZvlzEIZNm_CPdblkW2P2B-jM_UrTXiCRa9G0FlEOm3CPdblzXlIXX_WIi81..; _UP_F7E_8D_=i74i77uRTztAyuMdjhMg5P%2F9Z2BH6e82%2BQlPywU%2FyONiPxEnpGdMdNA19wBgmTqcFbwb8O576EdLHpvA%2Fhicy1HUTu2LBlCPom4qTWeMFqCgN55FQx3lIyu%2B1OsWIKjG1w4pOGWcZjvuo4m2jHof9eRj66wpeTPO4r7NBD%2F4wEE0IpjIBHWretgcndtvmRjOY20Scc1suKlfi7vZDl008Lzk4nczcPLP1rSgCMVm5ws2Z%2BPyTAVLm9UiDHFvPYOeAzBC2mqjnZNcOdSeBvi%2BK5asGKJO8g4hJESc0IAMGL4VOMEdD6uzZJXMxBnUatpyAHLu79tlMNqP8TGNMQXXgvSqK5ufzR58ZeivnehV0qE%2FWt1yDEDt%2BfWrmT4mVs6zZWXvqpzmoV3MeygIUCEakjmqUpi%2BMoCaK%2BK%2BYrCYUbg6F8u7yQFbh%2F0Q7RCSfK2U6tAXQttwc%2FtDK7HYGyvolg%3D%3D; _UP_D_=pc; __pus=2eb6089f6adc925c84990d92d4b0a357AATxYXvdaua3M0/wzmavzcpctOjpswIgFj4OguYz8B9IMjvHcs1+HKw3RsCkW2WmBDJ/LEWScwAuGGxAnflk2dtu; __kp=d4369090-a5c9-11ef-99ff-eb56513d6563; __kps=AAR+MKJzuB1EFV71vHDnDI2y; __ktd=TcmqKEcNVDspmHVa57dEoQ==; __uid=AAR+MKJzuB1EFV71vHDnDI2y; __puus=86201bdf280762f8bcd6c02277777688AARY+SEBfFv/axwceCz9klafzgU4S0AEgPaNqk8fboyefALJS+BHbebcYTj5ITB8+kBajqLtorClUo5nbBo1mLYutJT+YIicTnMVJmUsMUMvNgd7SU+xxHJIeX1b4H+VSgAX9X4RG+ZuHSH38CLg0RleO+v2jU6sbMuR6lVZwydpMV01vS2j8fSVSZrQkZ/Wki83K3r0idg59DC5Q+BNed92"`
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
