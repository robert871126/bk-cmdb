<template>
    <div class="edit-label-list">
        <div class="scrollbar-box">
            <div class="label-item" v-for="(label, index) in list" :key="index">
                <div class="label-key" :class="{ 'is-error': errors.has('key-' + index) }">
                    <input class="cmdb-form-input" type="text"
                        :data-vv-name="'key-' + index"
                        v-validate="getValidateRules(index, 'key')"
                        v-model="label.key"
                        :placeholder="$t('标签键')">
                    <p class="input-error">{{errors.first('key-' + index)}}</p>
                </div>
                <div class="label-value" :class="{ 'is-error': errors.has('value-' + index) }">
                    <input class="cmdb-form-input" type="text"
                        :data-vv-name="'value-' + index"
                        v-validate="getValidateRules(index, 'value')"
                        v-model="label.value"
                        :placeholder="$t('标签值')">
                    <p class="input-error">{{errors.first('value-' + index)}}</p>
                </div>
                <i class="bk-icon icon-plus-circle-shape icon-btn"
                    v-show="list.length - 1 === index"
                    @click="handleAddLabel(index)">
                </i>
                <i :class="['bk-icon', 'icon-minus-circle-shape', 'icon-btn']"
                    @click="handleRemoveLabel(index)">
                </i>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        props: {
            defaultList: {
                type: Array,
                default: () => []
            }
        },
        data () {
            return {
                originList: [],
                list: [],
                removeKeysList: [],
                submitList: []
            }
        },
        watch: {
            defaultList: {
                handler (list) {
                    let cloneList = this.$tools.clone(list)
                    if (!cloneList.length) {
                        cloneList = cloneList.concat([{
                            id: -1,
                            key: '',
                            value: ''
                        }])
                    }
                    this.list = cloneList
                    this.originList = this.$tools.clone(list)
                },
                deep: true
            },
            list: {
                handler (list) {
                    if (list.length === 1 && !list[0].key && !list[0].value) {
                        this.submitList = []
                        this.removeKeysList = this.originList.map(label => label.key)
                    } else {
                        const defaultList = this.defaultList
                        for (let i = 0; i < defaultList.length; i++) {
                            const index = list.findIndex(item => item.id !== -1 && defaultList[i].id === item.id && (defaultList[i].key !== item.key || defaultList[i].value !== item.value))
                            if (index !== -1) {
                                list[index].id = -1
                                !this.removeKeysList.includes[defaultList[i].key] && this.removeKeysList.push(defaultList[i].key)
                            }
                        }
                        this.submitList = list
                    }
                },
                deep: true
            }
        },
        created () {
            this.initValue()
        },
        methods: {
            initValue () {
                let cloneList = this.$tools.clone(this.defaultList)
                if (!cloneList.length) {
                    cloneList = cloneList.concat([{
                        id: -1,
                        key: '',
                        value: ''
                    }])
                }
                this.list = cloneList
                this.originList = this.$tools.clone(this.defaultList)
            },
            getValidateRules (currentIndex, type) {
                const rules = {}
                if (this.list.length === 1 && !this.list[0].key && !this.list[0].value) {
                    return {}
                }
                rules.required = true
                rules.instanceTag = true
                if (type === 'key') {
                    const list = this.$tools.clone(this.list)
                    list.splice(currentIndex, 1)
                    rules.repeatTagKey = list.map(label => label.key)
                }
                return rules
            },
            handleAddLabel (index) {
                const currentTag = this.list[index]
                if (!currentTag.key || !currentTag.value) {
                    this.$bkMessage({
                        message: this.$t('请填写完整标签键/值'),
                        theme: 'warning'
                    })
                    return
                }
                this.list.push({
                    id: -1,
                    key: '',
                    value: ''
                })
            },
            handleRemoveLabel (index) {
                if (this.list.length === 1) {
                    this.$set(this.list[0], 'key', '')
                    this.$set(this.list[0], 'value', '')
                    return
                }
                const currentTag = this.list[index]
                if (currentTag.id !== -1) {
                    this.removeKeysList.push(currentTag.key)
                }
                this.list.splice(index, 1)
            }
        }
    }
</script>

<style lang="scss" scoped>
    .edit-label-list {
        padding: 8px 12px 18px 26px;
        .scrollbar-box {
            @include scrollbar-y;
            height: 315px;
        }
        .label-item {
            display: flex;
            align-items: center;
            margin-bottom: 20px;
            &:last-child {
                margin-bottom: 10px;
            }
        }
        .cmdb-form-input {
            height: 32px;
            line-height: 32px;
        }
        .label-key {
            position: relative;
            width: 172px;
            margin-right: 10px;
        }
        .label-value {
            position: relative;
            width: 292px;
            margin-right: 10px;
        }
        .icon-btn {
            color: #c4c6cc;
            font-size: 18px;
            margin-right: 8px;
            cursor: pointer;
            &.disabled {
                color: #dcdee5;
                cursor: not-allowed;
                &:hover {
                    color: #dcdee5 !important;
                }
            }
            &:hover {
                color: #979ba5;
            }
        }
        .input-error {
            position: absolute;
            top: 100%;
            left: 0;
            line-height: 14px;
            font-size: 12px;
            color: #ff5656;
        }
        .is-error input.cmdb-form-input {
            border-color: #ff5656;
        }
    }
</style>
