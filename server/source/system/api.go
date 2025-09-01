package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i *initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwtAddedToBlackList"},

		{ApiGroup: "systemUser", Method: "DELETE", Path: "/user/deleteUser", Description: "deleteUsers"},
		{ApiGroup: "systemUser", Method: "POST", Path: "/user/admin_register", Description: "userRegistration"},
		{ApiGroup: "systemUser", Method: "POST", Path: "/user/getUserList", Description: "getUsersList"},
		{ApiGroup: "systemUser", Method: "PUT", Path: "/user/setUserInfo", Description: "setUserInfo"},
		{ApiGroup: "systemUser", Method: "PUT", Path: "/user/setSelfInfo", Description: "setSelfInfo"},
		{ApiGroup: "systemUser", Method: "GET", Path: "/user/getUserInfo", Description: "getSelfInfo"},
		{ApiGroup: "systemUser", Method: "POST", Path: "/user/setUserAuthorities", Description: "setPermissionGroup"},
		{ApiGroup: "systemUser", Method: "POST", Path: "/user/changePassword", Description: "changePassword"},
		{ApiGroup: "systemUser", Method: "POST", Path: "/user/setUserAuthority", Description: "modifyUserRole"},
		{ApiGroup: "systemUser", Method: "POST", Path: "/user/resetPassword", Description: "resetUserPassword"},
		{ApiGroup: "systemUser", Method: "PUT", Path: "/user/setSelfSetting", Description: "resetUserWeb"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "createAPI"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "deleteAPI"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "updateAPI"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "getAPIList"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "getAllAPI"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "getAPIByID"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "deleteAPIByID"},
		{ApiGroup: "api", Method: "GET", Path: "/api/syncApi", Description: "getSyncApi"},
		{ApiGroup: "api", Method: "GET", Path: "/api/getApiGroups", Description: "getRouteGroup"},
		{ApiGroup: "api", Method: "POST", Path: "/api/enterSyncApi", Description: "confirmSyncApi"},
		{ApiGroup: "api", Method: "POST", Path: "/api/ignoreApi", Description: "ignoreApi"},

		{ApiGroup: "role", Method: "POST", Path: "/authority/copyAuthority", Description: "copyRole"},
		{ApiGroup: "role", Method: "POST", Path: "/authority/createAuthority", Description: "createRole"},
		{ApiGroup: "role", Method: "POST", Path: "/authority/deleteAuthority", Description: "deleteRole"},
		{ApiGroup: "role", Method: "PUT", Path: "/authority/updateAuthority", Description: "updateRole"},
		{ApiGroup: "role", Method: "POST", Path: "/authority/getAuthorityList", Description: "getRoleList"},
		{ApiGroup: "role", Method: "POST", Path: "/authority/setDataAuthority", Description: "setRolePermissions"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "changeRoleAPIPermission"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "getPermissionList"},

		{ApiGroup: "menu", Method: "POST", Path: "/menu/addBaseMenu", Description: "addMenu"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/getMenu", Description: "getMenuTree"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "deleteMenu"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/updateBaseMenu", Description: "updateMenu"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/getBaseMenuById", Description: "getMenuByID"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/getMenuList", Description: "getMenuList"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "getDynamicRoute"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/getMenuAuthority", Description: "getMenuRole"},
		{ApiGroup: "menu", Method: "POST", Path: "/menu/addMenuAuthority", Description: "addMenuRole"},

		{ApiGroup: "partialUpload", Method: "POST", Path: "/fileUploadAndDownload/findFile", Description: "findTargetFile"},
		{ApiGroup: "partialUpload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "breakpointContinue"},
		{ApiGroup: "partialUpload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "breakpointContinueFinish"},
		{ApiGroup: "partialUpload", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "removeFileAfterUpload"},

		{ApiGroup: "fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "fileUploadExample"},
		{ApiGroup: "fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "deleteFile"},
		{ApiGroup: "fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "editFileNameOrRemark"},
		{ApiGroup: "fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "getUploadFileList"},
		{ApiGroup: "fileUploadDownload", Method: "POST", Path: "/fileUploadAndDownload/importURL", Description: "importURL"},

		{ApiGroup: "systemService", Method: "POST", Path: "/system/getServerInfo", Description: "getServerInfo"},
		{ApiGroup: "systemService", Method: "POST", Path: "/system/getSystemConfig", Description: "getConfigFileContent"},
		{ApiGroup: "systemService", Method: "POST", Path: "/system/setSystemConfig", Description: "setConfigFileContent"},

		{ApiGroup: "customer", Method: "PUT", Path: "/customer/customer", Description: "updateCustomer"},
		{ApiGroup: "customer", Method: "POST", Path: "/customer/customer", Description: "createCustomer"},
		{ApiGroup: "customer", Method: "DELETE", Path: "/customer/customer", Description: "deleteCustomer"},
		{ApiGroup: "customer", Method: "GET", Path: "/customer/customer", Description: "getSingleCustomer"},
		{ApiGroup: "customer", Method: "GET", Path: "/customer/customerList", Description: "getCustomerList"},

		{ApiGroup: "codeGenerator", Method: "GET", Path: "/autoCode/getDB", Description: "getAllDatabases"},
		{ApiGroup: "codeGenerator", Method: "GET", Path: "/autoCode/getTables", Description: "getDatabaseTables"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/createTemp", Description: "autoCode"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/preview", Description: "previewAutoCode"},
		{ApiGroup: "codeGenerator", Method: "GET", Path: "/autoCode/getColumn", Description: "getSelectedTableFields"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/installPlugin", Description: "installPlugin"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/pubPlug", Description: "packagePlugin"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/mcp", Description: "自动生成 MCP Tool 模板"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/mcpTest", Description: "MCP Tool 测试"},
		{ApiGroup: "codeGenerator", Method: "POST", Path: "/autoCode/mcpList", Description: "获取 MCP ToolList"},

		{ApiGroup: "templateConfiguration", Method: "POST", Path: "/autoCode/createPackage", Description: "configurationTemplates"},
		{ApiGroup: "templateConfiguration", Method: "GET", Path: "/autoCode/getTemplates", Description: "getTemplateFile"},
		{ApiGroup: "templateConfiguration", Method: "POST", Path: "/autoCode/getPackage", Description: "getAllTemplates"},
		{ApiGroup: "templateConfiguration", Method: "POST", Path: "/autoCode/delPackage", Description: "deleteTemplate"},

		{ApiGroup: "codeGenHistory", Method: "POST", Path: "/autoCode/getMeta", Description: "getMetaInfo"},
		{ApiGroup: "codeGenHistory", Method: "POST", Path: "/autoCode/rollback", Description: "rollbackAutoGeneratedCode"},
		{ApiGroup: "codeGenHistory", Method: "POST", Path: "/autoCode/getSysHistory", Description: "queryRollbackRecord"},
		{ApiGroup: "codeGenHistory", Method: "POST", Path: "/autoCode/delSysHistory", Description: "deleteRollbackRecord"},
		{ApiGroup: "codeGenHistory", Method: "POST", Path: "/autoCode/addFunc", Description: "addTemplateMethod"},

		{ApiGroup: "dictDetails", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "updateDictionaryContent"},
		{ApiGroup: "dictDetails", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "createDictionaryContent"},
		{ApiGroup: "dictDetails", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "deleteDictionaryContent"},
		{ApiGroup: "dictDetails", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "getDictionaryContentById"},
		{ApiGroup: "dictDetails", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "getDictionaryContentList"},

		{ApiGroup: "dictionary", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "createDictionary"},
		{ApiGroup: "dictionary", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "deleteDictionary"},
		{ApiGroup: "dictionary", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "updateDictionary"},
		{ApiGroup: "dictionary", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "getDictionaryById"},
		{ApiGroup: "dictionary", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "getDictionaryList"},

		{ApiGroup: "optRecord", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "createOperationRecord"},
		{ApiGroup: "optRecord", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "getOperationRecordById"},
		{ApiGroup: "optRecord", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "getOperationRecordList"},
		{ApiGroup: "optRecord", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "deleteOperationRecord"},
		{ApiGroup: "optRecord", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "batchDeleteOperationHistory"},

		{ApiGroup: "resumeUpload", Method: "POST", Path: "/simpleUploader/upload", Description: "pluginVersionResumableUpload"},
		{ApiGroup: "resumeUpload", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "fileIntegrityCheck"},
		{ApiGroup: "resumeUpload", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "mergeFileAfterUpload"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "sendTestEmail"},
		{ApiGroup: "email", Method: "POST", Path: "/email/sendEmail", Description: "sendEmail"},

		{ApiGroup: "buttonAuthority", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "setButtonPermission"},
		{ApiGroup: "buttonAuthority", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "getExistingButtonPermission"},
		{ApiGroup: "buttonAuthority", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "deleteButton"},

		{ApiGroup: "tableTemplate", Method: "POST", Path: "/sysExportTemplate/createSysExportTemplate", Description: "createExportTemplate"},
		{ApiGroup: "tableTemplate", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplate", Description: "deleteExportTemplate"},
		{ApiGroup: "tableTemplate", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplateByIds", Description: "batchDeleteExportTemplate"},
		{ApiGroup: "tableTemplate", Method: "PUT", Path: "/sysExportTemplate/updateSysExportTemplate", Description: "updateExportTemplate"},
		{ApiGroup: "tableTemplate", Method: "GET", Path: "/sysExportTemplate/findSysExportTemplate", Description: "getExportTemplateById"},
		{ApiGroup: "tableTemplate", Method: "GET", Path: "/sysExportTemplate/getSysExportTemplateList", Description: "getExportTemplateList"},
		{ApiGroup: "tableTemplate", Method: "GET", Path: "/sysExportTemplate/exportExcel", Description: "exportExcel"},
		{ApiGroup: "tableTemplate", Method: "GET", Path: "/sysExportTemplate/exportTemplate", Description: "downloadTemplate"},
		{ApiGroup: "tableTemplate", Method: "POST", Path: "/sysExportTemplate/importExcel", Description: "importExcel"},

		{ApiGroup: "announcement", Method: "POST", Path: "/info/createInfo", Description: "newAnnouncement"},
		{ApiGroup: "announcement", Method: "DELETE", Path: "/info/deleteInfo", Description: "deleteAnnouncement"},
		{ApiGroup: "announcement", Method: "DELETE", Path: "/info/deleteInfoByIds", Description: "batchDeleteAnnouncement"},
		{ApiGroup: "announcement", Method: "PUT", Path: "/info/updateInfo", Description: "updateAnnouncement"},
		{ApiGroup: "announcement", Method: "GET", Path: "/info/findInfo", Description: "getAnnouncementByID"},
		{ApiGroup: "announcement", Method: "GET", Path: "/info/getInfoList", Description: "getAnnouncementList"},

		{ApiGroup: "parameterManagement", Method: "POST", Path: "/sysParams/createSysParams", Description: "newParameter"},
		{ApiGroup: "parameterManagement", Method: "DELETE", Path: "/sysParams/deleteSysParams", Description: "deleteParameter"},
		{ApiGroup: "parameterManagement", Method: "DELETE", Path: "/sysParams/deleteSysParamsByIds", Description: "batchDeleteParameters"},
		{ApiGroup: "parameterManagement", Method: "PUT", Path: "/sysParams/updateSysParams", Description: "updateParameters"},
		{ApiGroup: "parameterManagement", Method: "GET", Path: "/sysParams/findSysParams", Description: "getParametersById"},
		{ApiGroup: "parameterManagement", Method: "GET", Path: "/sysParams/getSysParamsList", Description: "getParametersList"},
		{ApiGroup: "parameterManagement", Method: "GET", Path: "/sysParams/getSysParam", Description: "getParametersList"},
		{ApiGroup: "mediaLibraryCategories", Method: "GET", Path: "/attachmentCategory/getCategoryList", Description: "categoryList"},
		{ApiGroup: "mediaLibraryCategories", Method: "POST", Path: "/attachmentCategory/addCategory", Description: "addEditCategory"},
		{ApiGroup: "mediaLibraryCategories", Method: "POST", Path: "/attachmentCategory/deleteCategory", Description: "deleteCategory"},

		{ApiGroup: "versionControl", Method: "GET", Path: "/sysVersion/findSysVersion", Description: "getVersion"},
		{ApiGroup: "versionControl", Method: "GET", Path: "/sysVersion/getSysVersionList", Description: "getVersionList"},
		{ApiGroup: "versionControl", Method: "GET", Path: "/sysVersion/downloadVersionJson", Description: "downloadVersionJson"},
		{ApiGroup: "versionControl", Method: "POST", Path: "/sysVersion/exportVersion", Description: "createVersion"},
		{ApiGroup: "versionControl", Method: "POST", Path: "/sysVersion/importVersion", Description: "syncVersion"},
		{ApiGroup: "versionControl", Method: "DELETE", Path: "/sysVersion/deleteSysVersion", Description: "deleteVersion"},
		{ApiGroup: "versionControl", Method: "DELETE", Path: "/sysVersion/deleteSysVersionByIds", Description: "batchDeleteVersion"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+" "+global.Translate("general.tabelDataInitFail"))
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
